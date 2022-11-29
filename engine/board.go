package engine

type boardState struct {
	activeBoard
	players PlayerStates
	// Counts players that are not folded
	playersInRound int
	// Counts players that are not folded and are not AllIn
	actingPlayersInRound int
	// Counts players left in the game
	playersInGame int
}

func newBoardState() *boardState {
	return &boardState{
		activeBoard: activeBoard{
			communityCards: make([]*Card, 0, 5),
		},
		players: make([]*playerState, 8),
	}
}

// addCommunityCards adds community cards.
func (b *boardState) addCommunityCards(card ...*Card) {
	b.communityCards = append(b.communityCards, card...)
	b.iterateActivePlayers(func(p *playerState) {
		p.actor.SeeActiveBoardState(b.state())
	})
}

// seatPlayer finds an open seat for the player and puts him there.
func (b *boardState) seatPlayer(id string, bot BotPlayer) {
	for i, p := range b.players {
		if p == nil {
			b.players[i] = &playerState{
				activePlayerState: activePlayerState{
					seatNumber: i,
					id:         id,
				},
				actor: bot}
			return
		}
	}
}

// nextActiveSeat returns the index of the next seat which has a player in it which is still in the round.
func (b *boardState) nextActiveSeat(onSeat int) int {
	originalSeat := onSeat
	onSeat++
	for originalSeat != onSeat {
		if onSeat == 8 {
			onSeat = 0
		}
		if b.players[onSeat] != nil && b.players[onSeat].Status() == PlayerStatusPlaying {
			return onSeat
		}
		onSeat++
	}
	return originalSeat
}

// playerFolded lets the board know that a player has folded so it can keep track of how many players remain in the
// round.
func (b *boardState) playerFolded() {
	b.playersInRound--
	b.actingPlayersInRound--
}

// playerWentAllIn lets the board know that a player went all in so that we can track how many actingPlayers are left.
func (b *boardState) playerWentAllIn() {
	b.actingPlayersInRound--
}

// prevActiveSeat returns the index of the previous seat which has a player in it which is still in the round.
func (b *boardState) prevActiveSeat(onSeat int) int {
	originalSeat := onSeat
	onSeat--
	for originalSeat != onSeat {
		if onSeat == -1 {
			onSeat = 7
		}
		if b.players[onSeat] != nil && b.players[onSeat].Status() == PlayerStatusPlaying {
			return onSeat
		}
		onSeat--
	}
	return originalSeat
}

func (b *boardState) playerAtSeat(seat int) *playerState {
	return b.players[seat]
}

// moveSmallBlindButton moves the dealer button to the next non-empty seat.
func (b *boardState) moveSmallBlindButton() {
	b.smallBlindButton = b.nextActiveSeat(b.smallBlindButton)
}

func (b *boardState) nextStage() Stage {
	b.stage = b.stage.nextStage()
	b.iterateActivePlayers(func(p *playerState) {
		p.newStage()
	})
	return b.stage
}

// newRound performs the following actions
// 1. Call newRound on all players
// // 1. Clears out the players cards
// // 2. Marks the player as out of the game if their stack is at 0
// // Resets chipsEnteredThisStage and chipsEnteredThisRound
// 2. move the small blind button
// 3. reset the pot to 0
// 4. reset the stage to PreFlop
func (b *boardState) newRound() {
	b.playersInRound = 0
	b.iterateActivePlayers(func(p *playerState) {
		p.newRound()
		if p.status == PlayerStatusPlaying {
			b.playersInRound++
		}
	})
	b.actingPlayersInRound = b.playersInRound
	b.moveSmallBlindButton()
	b.pot = 0
	b.stage = PreFlop
}

// addToPot adds an amount to the pot.
func (b *boardState) addToPot(amount int) {
	b.pot += amount
}

func (b *boardState) iterateActivePlayers(f PlayerStateFunc) {
	b.playersInGame = 0
	b.players.iterateActive(b.smallBlindButton, func(p *playerState) {
		b.playersInGame++
		f(p)
	})
}

func (b *boardState) iterateActivePlayersFromTo(fromSeat, toSeat int, f PlayerStateFunc) {
	b.playersInGame = 0
	b.players.iterateActiveUntil(fromSeat, toSeat, func(p *playerState) {
		f(p)
		b.playersInGame = 0
	})
}

func (b *boardState) state() ActiveBoard {
	b.activeBoard.vPlayers = make([]ActivePlayerState, 8)
	for i, p := range b.players {
		b.activeBoard.vPlayers[i] = p.visiblePlayerState()
	}
	return &b.activeBoard
}
