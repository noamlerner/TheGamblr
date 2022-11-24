package pokerengine

type BoardPlayer struct {
	stack int
	// Chips entered on PreFlop, Flop, Turn or River.
	chipsEnteredThisStage int
	// Chips entered anytime since cards were received.
	chipsEnteredThisRound int
	cards                 Cards
	status                BoardPlayerStatus
	id                    string
	actor                 BotPlayer
	seatNumber            int
}

func (b *BoardPlayer) SetSeatNumber(seatNumber int) {
	b.seatNumber = seatNumber
}

func (b *BoardPlayer) Fold() {
	b.status = BoardPlayerStatusFolded
}

func (b *BoardPlayer) ReceiveCards(cards Cards, blindAmount int) int {
	b.chipsEnteredThisStage = 0
	b.cards = cards
	b.actor.ReceiveCards(cards, blindAmount)
	return b.RemoveChips(blindAmount)
}

func (b *BoardPlayer) Status() BoardPlayerStatus {
	return b.status
}

// NewStage indicates Flop, Turn or River
func (b *BoardPlayer) NewStage() {
	b.chipsEnteredThisStage = 0
}

// NewRound performs the following actions
// 1. Clears out the players cards
// 2. Marks the player as out of the game if their stack is at 0
// 3. Resets chipsEnteredThisStage and chipsEnteredThisRound
func (b *BoardPlayer) NewRound() {
	b.cards = nil
	if b.stack > 0 {
		b.status = BoardPlayerStatusPlaying
	} else {
		b.status = BoardPlayerStatusOut
	}
	b.chipsEnteredThisStage = 0
	b.chipsEnteredThisRound = 0
}

func (b *BoardPlayer) WinChips(amount int) {
	b.stack += amount
}

func (b *BoardPlayer) RemoveChips(amount int) int {
	if amount > b.stack {
		b.status = BoardPlayerStatusAllIn
		amount = b.stack
	}
	b.stack -= amount
	b.chipsEnteredThisStage += amount
	b.chipsEnteredThisRound += amount
	return amount
}
