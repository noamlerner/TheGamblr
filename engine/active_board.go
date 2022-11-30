package engine

type BoardState interface {
	// CommunityCards are the cards currently on the board. The first three cards will always be the flop, then the
	// turn, then the river.
	CommunityCards() Cards
	// Pot returns the size of the current Pot.
	Pot() int
	// Stage returns what the current Stage is, PreFlop, Flop, Turn or River.
	Stage() Stage
	// SmallBlindButton returns the index of the player in the Players slice that corresponds to the Small Blind Button.
	SmallBlindButton() int
	// Players return the array of all players in the game.
	Players() []PlayerState
}

// visibleBoardState is the board that all players can see during and at the start of a round.
type visibleBoardState struct {
	communityCards   []*Card
	pot              int
	stage            Stage
	smallBlindButton int
	vPlayers         []PlayerState
}

func (v *visibleBoardState) CommunityCards() Cards {
	return v.communityCards
}

func (v *visibleBoardState) Pot() int {
	return v.pot
}

func (v *visibleBoardState) Stage() Stage {
	return v.stage
}

func (v *visibleBoardState) SmallBlindButton() int {
	return v.smallBlindButton
}

func (v *visibleBoardState) Players() []PlayerState {
	return v.vPlayers
}
