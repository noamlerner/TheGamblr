package pokerengine

import (
	"sort"
)

type Dealer struct {
	deck       *Deck
	board      *Board
	gameConfig *GameConfig
	// any chips that were in a split pot but couldn't be divided evenly
	carryOverPot int
}

func NewDealer() *Dealer {
	return &Dealer{
		gameConfig: NewDefaultGameConfig(),
		deck:       NewDeck(),
		board:      NewBoard(),
	}
}

func (d *Dealer) PlayRound() HandedBoardPlayers {
	d.NewRound()
	// PreFlop Betting
	d.Betting()

	for i := 0; i < 3; i++ {
		winners := d.FindWinners()
		if winners != nil {
			return winners
		}
		d.NextStage()
	}

	return d.FindWinners()
}

func (d *Dealer) CashOutRound(winners HandedBoardPlayers) {
	if winners.Len() == 1 {
		// We have a winner!
		winners[0].p.WinChips(d.board.pot)
		return
	}

	// if true, then there is no split pot
	winningHand := winners[0].h.Beats(winners[1].h)
	if winningHand && winners[0].p.Status() != BoardPlayerStatusAllIn {
		// Winners[0] gets the whole pot
		winners[0].p.WinChips(d.board.pot)
		return
	} else if winningHand && winners[0].p.Status() == BoardPlayerStatusAllIn {
		d.splitPot([]*HandedBoardPlayer{winners[0]})
		if d.board.pot == 0 {
			return
		} else {
			d.CashOutRound(winners[1:])
			return
		}
	}

	numWinners := 2
	for i := 1; i < len(winners)-1; i++ {
		if winners[i].h.Beats(winners[i+1].h) {
			break
		}
		numWinners++
	}
	d.splitPot(winners[:numWinners])
	if d.board.pot == 0 {
		return
	}

	d.CashOutRound(winners[numWinners:])
}

// splitPot should only be called on all winners. We will split the pot between all passed in players.
func (d *Dealer) splitPot(winners HandedBoardPlayers) {
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
	d.board.IterateActivePlayers(func(p *BoardPlayer) {
		chipCount[p.id] = p.chipsEnteredThisRound
	})
	sort.Ints(chipCountsContributed)

	// now we figure out the size of the pots we are going to split
	pots := map[int]int{}
	d.board.IterateActivePlayers(func(p *BoardPlayer) {
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
				w.p.WinChips(winnings)
			}
		}

		carryOver := p - winnings*numPotParticipants
		d.carryOverPot += carryOver
		d.board.pot -= carryOver
	}
}

// NewRound performs the following actions
// 1. Shuffle the deck
// 2. Call board.NewRound
// // 1. Call NewRound on all players
// //// 1. Clears out the players cards
// //// 2. Marks the player as out of the game if their stack is at 0
// //// 3.Resets chipsEnteredThisStage and chipsEnteredThisRound
// // 2. move the small blind button
// // 3. reset the pot to 0
// // 4. reset the stage to PreFlop
// 3. Deals the cards and assigns the blinds
func (d *Dealer) NewRound() {
	d.deck.Shuffle()
	onPlayer := 0
	d.board.NewRound()
	d.board.IterateActivePlayers(func(p *BoardPlayer) {
		if p.Status() != BoardPlayerStatusPlaying {
			return
		}
		addedToPot := 0
		switch onPlayer {
		case 0:
			// smallBlind
			addedToPot = p.ReceiveCards(d.deck.NextCards(2), d.gameConfig.smallBlind)
		case 1:
			// big blind
			addedToPot = p.ReceiveCards(d.deck.NextCards(2), d.gameConfig.smallBlind*2)
		default:
			addedToPot = p.ReceiveCards(d.deck.NextCards(2), 0)
		}
		onPlayer++
		d.board.AddToPot(addedToPot)
	})
	d.board.AddToPot(d.carryOverPot)
	d.carryOverPot = 0
}

func (d *Dealer) NextStage() {
	switch d.board.NextStage() {
	case Flop:
		d.board.AddCommunityCards(d.deck.NextCards(3)...)
	case Turn:
		d.board.AddCommunityCards(d.deck.NextCard())
	case River:
		d.board.AddCommunityCards(d.deck.NextCard())
	}
}
func (d *Dealer) Betting() {
	// First To go
	startAt := d.board.smallBlindButton
	// Last to go exclusive
	endAt := d.board.smallBlindButton
	iterate := true
	// amount we need to call
	callAmount := 0
	// A raise must move the callAmount to this value
	minRaiseTo := d.gameConfig.SmallBlind() * 2

	// start/end positions, callAmount and raise to are different preflop
	if d.board.stage == PreFlop {
		callAmount = d.gameConfig.smallBlind * 2
		minRaiseTo = callAmount * 2
		startAt = d.board.NextActiveSeat(d.board.NextActiveSeat(d.board.smallBlindButton))
		endAt = startAt
	}

	for iterate == true {
		iterate = false
		d.board.IterateActivePlayersFromTo(startAt, endAt, func(player *BoardPlayer) {
			if player.Status() == BoardPlayerStatusAllIn ||
				player.Status() == BoardPlayerStatusFolded ||
				d.board.actingPlayersInRound == 1 {
				// We do not want to play folded players.
				// We also skip everyone if there is only one player left - that player won.
				return
			}
			action, raiseTo := player.actor.Act()
			if action == RaiseAction && player.stack == 0 {
				action = CallAction
			}
			switch action {
			case CheckFold:
				if callAmount-player.chipsEnteredThisStage > 0 {
					player.Fold()
				}
			case FoldAction:
				// Player Folds, we will move on to the next player.
				player.Fold()
			case CallAction:
				// player calls, this might put them all in automatically.
				d.board.AddToPot(player.RemoveChips(callAmount - player.chipsEnteredThisStage))
			case RaiseAction:
				// Player Raises. There is a MinRaise, if they tried to Raise less than that, we
				// will automatically put them there. The first MinRaise == BigBlind.
				if raiseTo < minRaiseTo {
					raiseTo = minRaiseTo
				}
				chipsToEnter := player.RemoveChips(raiseTo - player.chipsEnteredThisStage)
				// In case this puts them all in, we need to recalculate the actual new raiseTo
				raiseTo = player.chipsEnteredThisStage
				// Next raise must be at least the amount of this raise
				minRaiseTo = (raiseTo - callAmount) + raiseTo
				callAmount = raiseTo
				// Add chips into the board from the player.
				d.board.AddToPot(chipsToEnter)

				if endAt != player.seatNumber {
					// This is an edge case that can only happen once per stage - if the player to raise is also the
					// first to act and that player raised on their first action. In that case, we don't get
					// an extra iteration.
					startAt = endAt
					endAt = player.seatNumber
					iterate = true
				}
			}
			if player.Status() == BoardPlayerStatusFolded {
				d.board.PlayerFolded()
			}
			if player.Status() == BoardPlayerStatusAllIn {
				d.board.PlayerWentAllIn()
			}
		})
	}
}

// FindWinners will return nil if the game is not eligible to be won yet.
// It is assumed that whatever stage d.board is in has just ended.
// If everyone folded, we will return the last one left
// otherwise, we will return all remaining players in order of strongest to weakest.
func (d *Dealer) FindWinners() HandedBoardPlayers {
	if d.board.playersInRound == 1 {
		var lastManStanding *BoardPlayer
		d.board.IterateActivePlayers(func(p *BoardPlayer) {
			if p.Status() == BoardPlayerStatusPlaying {
				lastManStanding = p
			}
		})
		return []*HandedBoardPlayer{
			{
				h: NewHand(append(d.board.communityCards, lastManStanding.cards...)),
				p: lastManStanding,
			},
		}
	}

	if d.board.stage != River {
		// no winner yet
		return nil
	}

	handedBoardPlayers := make(HandedBoardPlayers, 0, d.board.playersInRound)
	d.board.IterateActivePlayers(func(p *BoardPlayer) {
		if !p.Status().EligibleToWin() {
			return
		}
		h := NewHand(append(d.board.communityCards, p.cards...))
		handedBoardPlayers = append(handedBoardPlayers, &HandedBoardPlayer{p: p, h: h})
	})
	sort.Sort(handedBoardPlayers)

	return handedBoardPlayers
}
