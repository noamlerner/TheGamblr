package pokerengine

type Board struct {
	communityCards []*Card
	pot            int
	stage          Stage

	players BoardPlayers
	// Counts players that are not folded
	playersInRound int
	// Counts players that are not folded and are not AllIn
	actingPlayersInRound int
	smallBlindButton     int
}

func NewBoard() *Board {
	return &Board{
		communityCards: make([]*Card, 0, 5),
		players:        make([]*BoardPlayer, 8),
	}
}

// AddCommunityCards adds community cards.
func (b *Board) AddCommunityCards(card ...*Card) {
	b.communityCards = append(b.communityCards, card...)
}

// Flop returns the current flop (first 3 community cards).
func (b *Board) Flop() []*Card {
	return b.communityCards[:3]
}

// Turn returns the current Turn (4th community card).
func (b *Board) Turn() *Card {
	return b.communityCards[3]
}

// River returns the current River (5th community card).
func (b *Board) River() *Card {
	return b.communityCards[4]
}

// SeatPlayer finds an open seat for the player and puts him there.
func (b *Board) SeatPlayer(id string, bot BotPlayer) {
	for i, p := range b.players {
		if p == nil {
			b.players[i] = &BoardPlayer{id: id, actor: bot, seatNumber: i}
			return
		}
	}
}

// NextActiveSeat returns the index of the next seat which has a player in it which is still in the round.
func (b *Board) NextActiveSeat(onSeat int) int {
	originalSeat := onSeat
	onSeat++
	for originalSeat != onSeat {
		if onSeat == 8 {
			onSeat = 0
		}
		if b.players[onSeat] != nil && b.players[onSeat].Status() == BoardPlayerStatusPlaying {
			return onSeat
		}
		onSeat++
	}
	return originalSeat
}

// PlayerFolded lets the board know that a player has folded so it can keep track of how many players remain in the
// round.
func (b *Board) PlayerFolded() {
	b.playersInRound--
	b.actingPlayersInRound--
}

// PlayerWentAllIn lets the board know that a player went all in so that we can track how many actingPlayers are left.
func (b *Board) PlayerWentAllIn() {
	b.actingPlayersInRound--
}

// PrevActiveSeat returns the index of the previous seat which has a player in it which is still in the round.
func (b *Board) PrevActiveSeat(onSeat int) int {
	originalSeat := onSeat
	onSeat--
	for originalSeat != onSeat {
		if onSeat == -1 {
			onSeat = 7
		}
		if b.players[onSeat] != nil && b.players[onSeat].Status() == BoardPlayerStatusPlaying {
			return onSeat
		}
		onSeat--
	}
	return originalSeat
}

func (b *Board) PlayerAtSeat(seat int) *BoardPlayer {
	return b.players[seat]
}

// moveSmallBlindButton moves the dealer button to the next non-empty seat.
func (b *Board) moveSmallBlindButton() {
	b.smallBlindButton = b.NextActiveSeat(b.smallBlindButton)
}

func (b *Board) NextStage() Stage {
	b.stage = b.stage.NextStage()
	b.IterateActivePlayers(func(p *BoardPlayer) {
		p.NewStage()
	})
	return b.stage
}

// NewRound performs the following actions
// 1. Call NewRound on all players
// // 1. Clears out the players cards
// // 2. Marks the player as out of the game if their stack is at 0
// // Resets chipsEnteredThisStage and chipsEnteredThisRound
// 2. move the small blind button
// 3. reset the pot to 0
// 4. reset the stage to PreFlop
func (b *Board) NewRound() {
	b.playersInRound = 0
	b.IterateActivePlayers(func(p *BoardPlayer) {
		p.NewRound()
		if p.status == BoardPlayerStatusPlaying {
			b.playersInRound++
		}
	})
	b.actingPlayersInRound = b.playersInRound
	b.moveSmallBlindButton()
	b.pot = 0
	b.stage = PreFlop
}

// AddToPot adds an amount to the pot.
func (b *Board) AddToPot(amount int) {
	b.pot += amount
}
func (b *Board) IterateActivePlayers(f BoardPlayerFunc) {
	b.players.IterateActive(b.smallBlindButton, f)
}

func (b *Board) IterateActivePlayersFromTo(fromSeat, toSeat int, f BoardPlayerFunc) {
	b.players.IterateActiveUntil(fromSeat, toSeat, f)
}
