package engine

import "math/rand"

// BotPlayerProvider returns a BotPlayer. We want a provider so that the caller can choose to reset the player state
// or not reset it between games.
// They also need to return a name for their bot. These names must be unique for accurate results.
type BotPlayerProvider func() (string, BotPlayer)

// GameSim can be used to measure how good your bot is. You can set a game config and how many games you want the bots
// to play. We will then run the game that many times and report back the results.
type GameSim struct {
	// PlayerProviders Read the comment above BotPlayerProvider
	PlayerProviders []BotPlayerProvider

	// NumSims - how many games are we going to run?
	NumSims int
	// GameConfig - what game config are we going to use for each simulation?
	GameConfig *GameConfig
}

func NewGameSim(playerProviders []BotPlayerProvider) *GameSim {
	return &GameSim{PlayerProviders: playerProviders, NumSims: 500, GameConfig: NewDefaultGameConfig()}
}

func (g *GameSim) WithNumSims(numSims int) *GameSim {
	g.NumSims = numSims
	return g
}

func (g *GameSim) WithGameConfig(config *GameConfig) *GameSim {
	g.GameConfig = config
	return g
}

// Run runs the game sim. Since games can end with multiple players having chips, every user will get of a win based
// on how many chips they have at the end of the round.
// We normalize for number of simulations.
func (g *GameSim) Run() map[string]float64 {
	results := make(map[string]float64, len(g.PlayerProviders))
	for _, pp := range g.PlayerProviders {
		n, _ := pp()
		results[n] = 0
	}

	totalChipsOnTable := float64(g.GameConfig.StartingStack * len(g.PlayerProviders))
	for i := 0; i < g.NumSims; i++ {
		dealer := NewDealer(g.GameConfig)
		rand.Shuffle(len(g.PlayerProviders), func(i, j int) {
			g.PlayerProviders[i], g.PlayerProviders[j] = g.PlayerProviders[j], g.PlayerProviders[i]
		})

		for _, pp := range g.PlayerProviders {
			n, p := pp()
			dealer.SeatPlayer(n, p)
		}
		finalBoardState := dealer.RunGame()
		for _, p := range finalBoardState.Players() {
			if p == nil {
				continue
			}
			results[p.Id()] += (float64(p.Stack()) / totalChipsOnTable) / float64(g.NumSims)
		}
	}
	return results
}
