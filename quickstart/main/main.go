package main

import (
	"github.com/noamlerner/TheGamblr/engine"
	"github.com/noamlerner/TheGamblr/quickstart"
)

func main() {
	playerProviders := []engine.BotPlayerProvider{
		func() (string, engine.BotPlayer) {
			return "TheGamblr", engine.NewRandomActionBot()
		},
		func() (string, engine.BotPlayer) {
			return "Bean", quickstart.NewOurBot()
		},
	}
	gameSim := engine.NewGameSim(playerProviders).WithNumSims(1)

	gameConfig := &engine.GameConfig{
		// SmallBlind will be 5
		SmallBlind: 5,
		// They only have 10 rounds until the game is over
		NumRounds: 10,
		// each player starts with 100 chips
		StartingStack: 100,
		// The highest log level, everything will be outputted.
		LogLevel: engine.LogLevelCards,
	}

	gameSim.WithGameConfig(gameConfig).Run()
}
