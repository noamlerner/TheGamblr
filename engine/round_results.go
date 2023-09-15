package engine

type roundEndBoard struct {
	activeBoard
	roundEndPlayers []PlayerResults
}

func (r *roundEndBoard) PlayerResults() []PlayerResults {
	return r.roundEndPlayers
}

type RoundResults interface {
	// CommunityCards are the cards currently on the board. The first three cards will always be the flop, then the
	// turn, then the river.
	CommunityCards() []*Card
	// Pot returns the size of the current Pot.
	Pot() int
	// SmallBlindButton returns the index of the player in the Players slice that corresponds to the Small Blind Button.
	SmallBlindButton() int
	// PlayerResults return the array of all players in the game.
	PlayerResults() []PlayerResults
}
