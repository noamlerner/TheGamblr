package engine

import (
	"sort"
)

type Dealer struct {
	deck        Deck
	board       *boardState
	gameConfig  *GameConfig
	lastToRaise *playerState
	log         *logger
	// any chips that were in a split pot but couldn't be divided evenly
	carryOverPot int
}

func DealerWithDefaultConfig() *Dealer {
	return NewDealer(NewDefaultGameConfig())
}

func NewDealer(config *GameConfig) *Dealer {
	return &Dealer{
		gameConfig: config,
		deck:       NewDeck(),
		board:      newBoardState(),
		log:        &logger{logLevel: config.LogLevel},
	}
}

func (d *Dealer) RunGame() BoardState {

	if d.gameConfig.NumRounds == -1 {
		roundNum := 0
		for d.board.playersInGame > 1 {
			d.log.Round(roundNum)
			d.playRound()
			roundNum++
		}

	} else {
		for i := 0; i < d.gameConfig.NumRounds && d.board.playersInGame > 1; i++ {
			d.log.Round(i)
			d.playRound()
		}
	}

	state := d.board.state()
	state.(*visibleBoardState).stage = GameOver
	d.board.iterateActivePlayers(func(p *playerState) {
		p.actor.SeeBoardState(state)
	})
	return state
}

// SeatPlayer return the seat number
func (d *Dealer) SeatPlayer(id string, player BotPlayer) int {
	d.board.playersInGame++
	return d.board.seatPlayer(id, player, d.gameConfig.StartingStack)
}

func (d *Dealer) GameConfig() *GameConfig {
	return d.gameConfig
}

func (d *Dealer) playRound() EndRoundPlayers {
	d.newRound()
	// PreFlop betting
	d.betting()

	for i := 0; i < 3; i++ {
		winners := d.findWinners()
		if winners != nil {
			d.endRound(winners)
			return winners
		}
		d.nextStage()
		d.betting()
	}

	winners := d.findWinners()
	d.endRound(winners)
	return winners
}

func (d *Dealer) endRound(endRoundPlayers EndRoundPlayers) {
	d.cashOutRound(endRoundPlayers)
	winnersByID := d.whoShowsTheirHand(endRoundPlayers)

	state := d.board.state()
	d.board.iterateActivePlayers(func(p *playerState) {
		w, ok := winnersByID[p.id]
		if !ok {
			w = &playerForWinnerCalculations{
				p: p,
			}
		}
		w.p.roundEndStats = w.toState()
		state.Players()[w.p.seatNumber] = w.p
	})
	state.(*visibleBoardState).stage = RoundOver
	d.log.Winners(state)
	d.board.iterateActivePlayers(func(p *playerState) {
		p.actor.SeeBoardState(state)
	})
}

// whoShowsTheirHand takes the EndRoundPlayers after cashOutRound is complete and sets willShowHand to true on anyone
// required to show their hand.
func (d *Dealer) whoShowsTheirHand(winners EndRoundPlayers) map[string]*playerForWinnerCalculations {
	if len(winners) == 1 {
		return map[string]*playerForWinnerCalculations{
			winners[0].p.id: winners[0],
		}
	}
	winnersByID := make(map[string]*playerForWinnerCalculations, len(winners))
	playerIdToHandRank := make(map[string]int, len(winners))
	for i, w := range winners {
		if w.chipsWon != 0 {
			// gotta show your hand to get chips
			w.willShowHand = true
		}
		winnersByID[w.p.id] = w
		// players are sorted in order of hand strength. Strongest hand comes first and so on.
		playerIdToHandRank[w.p.id] = i
	}

	strongestHand := -1
	d.board.iterateActivePlayersFromTo(d.lastToRaise.seatNumber, d.lastToRaise.seatNumber, func(p *playerState) {
		w, ok := winnersByID[p.id]
		if !ok {
			// This player did not make it to the final ronud.
			return
		}
		if w.willShowHand {
			strongestHand = playerIdToHandRank[p.id]
			// we already know they will show their hand
			return
		}
		// last to raise always shows hand
		if strongestHand == -1 {
			w.willShowHand = true
			strongestHand = playerIdToHandRank[p.id]
			return
		}

		handRank := playerIdToHandRank[p.id]
		if handRank < strongestHand {
			// strongest hand so far, so this player won't muck
			w.willShowHand = true
			strongestHand = handRank
		}
	})
	return winnersByID
}

func (d *Dealer) cashOutRound(winners EndRoundPlayers) {
	if winners.Len() == 0 {
		return
	}
	if winners.Len() == 1 {
		// We have a winner!
		winners[0].winChips(d.board.pot)
		return
	}

	// if true, then there is no split pot
	winningHand := winners[0].hand.Beats(winners[1].hand)
	if winningHand && winners[0].p.Status() != PlayerStatusAllIn {
		// LogLevelWinners[0] gets the whole pot
		winners[0].winChips(d.board.pot)
		return
	} else if winningHand && winners[0].p.Status() == PlayerStatusAllIn {
		d.splitPot([]*playerForWinnerCalculations{winners[0]})
		if d.board.pot == 0 {
			return
		} else {
			d.cashOutRound(winners[1:])
			return
		}
	}

	numWinners := 2
	for i := 1; i < len(winners)-1; i++ {
		if winners[i].hand.Beats(winners[i+1].hand) {
			break
		}
		numWinners++
	}
	d.splitPot(winners[:numWinners])
	if d.board.pot == 0 {
		return
	}

	d.cashOutRound(winners[numWinners:])
}

// splitPot should only be called on all winners. We will split the pot between all passed in players.
func (d *Dealer) splitPot(winners EndRoundPlayers) {
	// we start by figuring out how many pots there are. This should be 1 unless people went all in.
	seenChipCount := map[int]bool{}
	// we will iterate over this in order to build the different pots that can be split.
	chipCountsContributed := []int{}
	for _, w := range winners {
		if _, ok := seenChipCount[w.p.chipsEnteredThisRound]; !ok {
			chipCountsContributed = append(chipCountsContributed, w.p.chipsEnteredThisRound)
		}
		seenChipCount[w.p.chipsEnteredThisRound] = true
	}
	chipCount := map[string]int{}
	d.board.iterateActivePlayers(func(p *playerState) {
		chipCount[p.id] = p.chipsEnteredThisRound
	})
	sort.Ints(chipCountsContributed)

	// now we figure out the size of the pots we are going to split
	pots := map[int]int{}
	d.board.iterateActivePlayers(func(p *playerState) {
		for _, c := range chipCountsContributed {
			chipsToPot := MinInt(p.chipsEnteredThisRound, c)
			pots[c] += chipsToPot
			p.chipsEnteredThisRound -= chipsToPot
		}
	})

	// c is how much you have to have contributed to get from this pot, p is how much the pot actually is.
	for c, p := range pots {
		// how many ways are we splitting the pot?
		numPotParticipants := 0
		for _, w := range winners {
			if chipCount[w.p.id] >= c {
				numPotParticipants++
			}
		}

		winnings := p / numPotParticipants
		// split the pot!
		for _, w := range winners {
			if chipCount[w.p.id] >= c {
				d.board.pot -= winnings
				w.winChips(winnings)
			}
		}

		carryOver := p - winnings*numPotParticipants
		d.carryOverPot += carryOver
		d.board.pot -= carryOver
	}
}

// newRound performs the following actions
// 1. Shuffle the deck
// 2. Call board.newRound
// // 1. Call newRound on all players
// //// 1. Clears out the players cards
// //// 2. Marks the player as out of the game if their stack is at 0
// //// 3.Resets chipsEnteredThisStage and chipsEnteredThisRound
// // 2. move the small blind button
// // 3. reset the pot to 0
// // 4. reset the stage to PreFlop
// 3. Deals the cards and assigns the blinds
func (d *Dealer) newRound() {
	d.deck.Shuffle()
	onPlayer := 0
	d.board.newRound()
	d.board.iterateActivePlayers(func(p *playerState) {
		if p.Status() != PlayerStatusPlaying {
			return
		}
		switch onPlayer {
		case 0:
			// SmallBlind
			addedToPot := p.receiveCards(d.deck.NextCards(2), d.gameConfig.SmallBlind)
			d.board.addToPot(addedToPot)
			d.announceAction(p, SmallBlind, d.gameConfig.SmallBlind)
			d.lastToRaise = p
		case 1:
			// big blind
			addedToPot := p.receiveCards(d.deck.NextCards(2), d.gameConfig.SmallBlind*2)
			d.board.addToPot(addedToPot)
			d.announceAction(p, BigBlind, d.gameConfig.SmallBlind*2)
		default:
			p.receiveCards(d.deck.NextCards(2), 0)
		}
		onPlayer++
		d.log.Cards(p)
	})
	d.board.addToPot(d.carryOverPot)
	d.carryOverPot = 0
}

func (d *Dealer) announceAction(p *playerState, action ActionType, amount int) {
	vAction := newVisibleAction(p.visibleState(), action, amount)
	d.log.Action(vAction)
	state := d.board.state()
	d.board.iterateActivePlayers(func(p *playerState) {
		p.actor.ActionUpdate(vAction, state)
	})
}

func (d *Dealer) nextStage() {
	switch d.board.nextStage() {
	case Flop:
		d.board.addCommunityCards(d.deck.NextCards(3)...)
	case Turn:
		d.board.addCommunityCards(d.deck.NextCard())
	case River:
		d.board.addCommunityCards(d.deck.NextCard())
	}
	d.announceBoardState()
}

func (d *Dealer) announceBoardState() {
	state := d.board.state()
	d.log.Stage(state)
	d.board.iterateActivePlayers(func(p *playerState) {
		p.actor.SeeBoardState(state)
	})
}

func (d *Dealer) betting() {
	// First To go
	startAt := d.board.smallBlindButton
	// Last to go exclusive
	endAt := d.board.smallBlindButton
	iterate := true
	// amount we need to call
	callAmount := 0
	// A raise must move the callAmount to this value
	minRaiseTo := d.gameConfig.SmallBlind * 2

	// start/end positions, callAmount and raise to are different preflop
	if d.board.stage == PreFlop {
		callAmount = d.gameConfig.SmallBlind * 2
		minRaiseTo = callAmount * 2
		startAt = d.board.nextActiveSeat(d.board.nextActiveSeat(d.board.smallBlindButton))
		endAt = startAt
	}

	// will only be true for the first acting player on their first action
	firstAction := true
	// when someone raises, we need to loop again. We will always start at the person we just ended at.
	nextStartAt := endAt
	for iterate == true {
		iterate = false
		d.board.iterateActivePlayersFromTo(startAt, endAt, func(player *playerState) {
			if player.Status() == PlayerStatusAllIn ||
				player.Status() == PlayerStatusFolded ||
				d.board.actingPlayersInRound == 1 {
				// We do not want to play folded players.
				// We also skip everyone if there is only one player left - that player won.
				return
			}
			action, raiseTo := player.actor.Act(d.board.pot, callAmount, callAmount-player.chipsEnteredThisStage)
			if action == RaiseAction && player.stack == 0 {
				action = CallAction
			}
			switch action {
			case CheckFoldAction:
				if callAmount-player.chipsEnteredThisStage > 0 {
					player.fold()
					d.announceAction(player, FoldAction, 0)
				}
			case FoldAction:
				// Player Folds, we will move on to the next player.
				player.fold()
				d.announceAction(player, FoldAction, 0)
			case CallAction:
				// player calls, this might put them all in automatically.
				chipsEntered := player.removeChips(callAmount - player.chipsEnteredThisStage)
				d.board.addToPot(chipsEntered)
				if chipsEntered == 0 {
					d.announceAction(player, CheckFoldAction, chipsEntered)
				} else {
					d.announceAction(player, CallAction, chipsEntered)
				}
			case RaiseAction:
				// Player Raises. There is a MinRaise, if they tried to Raise less than that, we
				// will automatically put them there. The first MinRaise == BigBlind.
				if raiseTo < minRaiseTo {
					raiseTo = minRaiseTo
				}
				chipsToEnter := player.removeChips(raiseTo - player.chipsEnteredThisStage)
				// In case this puts them all in, we need to recalculate the actual new raiseTo
				raiseTo = player.chipsEnteredThisStage
				// Next raise must be at least the amount of this raise
				minRaiseTo = (raiseTo - callAmount) + raiseTo
				callAmount = raiseTo
				// Add chips into the board from the player.
				d.board.addToPot(chipsToEnter)
				d.announceAction(player, RaiseAction, callAmount)
				d.lastToRaise = player

				if !firstAction {
					// This is an edge case that can only happen once per stage - if the player to raise is also the
					// first to act and that player raised on their first action. In that case, we don't get
					// an extra iteration.
					endAt = player.seatNumber
					iterate = true
				}
			}
			if player.Status() == PlayerStatusFolded {
				d.board.playerFolded()
			}
			if player.Status() == PlayerStatusAllIn {
				d.board.playerWentAllIn()
			}
			firstAction = false
		})
		startAt = nextStartAt
		nextStartAt = endAt
	}
}

// findWinners will return nil if the game is not eligible to be won yet.
// It is assumed that whatever stage d.board is in has just ended.
// If everyone folded, we will return the last one left
// otherwise, we will return all remaining players in order of strongest to weakest.
func (d *Dealer) findWinners() EndRoundPlayers {
	if d.board.playersInRound == 1 {
		var lastManStanding *playerState
		d.board.iterateActivePlayers(func(p *playerState) {
			if p.Status() == PlayerStatusPlaying {
				lastManStanding = p
			}
		})
		return []*playerForWinnerCalculations{
			{
				hand:         NewHand(append(d.board.communityCards, lastManStanding.cards...)),
				p:            lastManStanding,
				willShowHand: false,
			},
		}
	}

	if d.board.stage != River {
		// no winner yet
		return nil
	}

	handedBoardPlayers := make(EndRoundPlayers, 0, d.board.playersInRound)
	d.board.iterateActivePlayers(func(p *playerState) {
		if !p.Status().eligibleToWin() {
			return
		}
		h := NewHand(append(d.board.communityCards, p.cards...))
		handedBoardPlayers = append(handedBoardPlayers, &playerForWinnerCalculations{p: p, hand: h})
	})
	sort.Sort(handedBoardPlayers)

	return handedBoardPlayers
}
