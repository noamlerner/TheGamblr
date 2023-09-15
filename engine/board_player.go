package engine

type playerState struct {
	visiblePlayerState
	// Chips entered on PreFlop, Flop, Turn or River.
	chipsEnteredThisStage int
	// Chips entered anytime since cards were received.
	chipsEnteredThisRound int
	cards                 Cards
	actor                 BotPlayer
}

func (b *playerState) setSeatNumber(seatNumber int) {
	b.seatNumber = seatNumber
}

func (b *playerState) fold() {
	b.status = PlayerStatusFolded
}

func (b *playerState) receiveCards(cards Cards, blindAmount int) int {
	b.chipsEnteredThisStage = 0
	b.cards = cards
	b.actor.ReceiveCards(cards)
	return b.removeChips(blindAmount)
}

func (b *playerState) Status() PlayerStatus {
	return b.status
}

// newStage indicates Flop, Turn or River
func (b *playerState) newStage() {
	b.chipsEnteredThisStage = 0
}

// newRound performs the following sequence
// 1. Clears out the players cards
// 2. Marks the player as out of the game if their stack is at 0
// 3. Resets chipsEnteredThisStage and chipsEnteredThisRound
func (b *playerState) newRound() {
	b.cards = nil
	if b.stack > 0 {
		b.status = PlayerStatusPlaying
	} else {
		b.status = PlayerStatusOut
	}
	b.chipsEnteredThisStage = 0
	b.chipsEnteredThisRound = 0
}

func (b *playerState) winChips(amount int) {
	b.stack += amount
}

func (b *playerState) removeChips(amount int) int {
	if amount >= b.stack {
		b.status = PlayerStatusAllIn
		amount = b.stack
	}
	b.stack -= amount
	b.chipsEnteredThisStage += amount
	b.chipsEnteredThisRound += amount
	return amount
}

func (b *playerState) visibleState() *visiblePlayerState {
	if b == nil {
		return nil
	}
	return &b.visiblePlayerState
}
