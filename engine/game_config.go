package engine

type GameConfig struct {
	// SmallBlind defines the games small blind. BigBlind will be x2.
	SmallBlind int
	// NumRounds defines how many rounds will be played in this game. If this is -1, we will play until there is only
	// one player left with chips.
	NumRounds int
}

func NewDefaultGameConfig() *GameConfig {
	return &GameConfig{
		SmallBlind: 5,
		NumRounds:  200,
	}
}
