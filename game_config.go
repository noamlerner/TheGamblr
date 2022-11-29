package TheGamblr

type GameConfig struct {
	smallBlind int
}

func NewDefaultGameConfig() *GameConfig {
	return &GameConfig{
		smallBlind: 5,
	}
}
func (g *GameConfig) SmallBlind() int {
	return g.smallBlind
}
