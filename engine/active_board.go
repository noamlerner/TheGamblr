package engine

type ActiveBoard interface {
	// CommunityCards are the cards currently on the board. The first three cards will always be the flop, then the
	// turn, then the river.
	CommunityCards() []*Card
	// Pot returns the size of the current Pot.
	Pot() int
	// Stage returns what the current Stage is, PreFlop, Flop, Turn or River.
	Stage() Stage
	// SmallBlindButton returns the index of the player in the Players slice that corresponds to the Small Blind Button.
	SmallBlindButton() int
	// Players return the array of all players in the game.
	Players() []ActivePlayerState
}

// activeBoard is the board that all players can see during and at the start of a round.
type activeBoard struct {
	communityCards   []*Card
	pot              int
	stage            Stage
	smallBlindButton int
	vPlayers         []ActivePlayerState
}

func (v *activeBoard) CommunityCards() []*Card {
	return v.communityCards
}

func (v *activeBoard) Pot() int {
	return v.pot
}

func (v *activeBoard) Stage() Stage {
	return v.stage
}

func (v *activeBoard) SmallBlindButton() int {
	return v.smallBlindButton
}

func (v *activeBoard) Players() []ActivePlayerState {
	return v.vPlayers
}
