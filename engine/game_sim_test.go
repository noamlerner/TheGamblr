package engine

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameSim(t *testing.T) {
	playerProviders := []BotPlayerProvider{
		func() (string, BotPlayer) {
			return "Noam", NewRandomActionBot()
		},
		func() (string, BotPlayer) {
			return "Bean", NewRandomActionBot()
		},
		func() (string, BotPlayer) {
			return "TheGamblr", NewRandomActionBot()
		},
	}
	results := NewGameSim(playerProviders).WithNumSims(1000).WithGameConfig(NewDefaultGameConfig()).Run()
	assert.NotNil(t, results)
	// At 1000 sims, we can trust that we are approaching the true number. They are the same bot so each has a 33%
	// chance of winning. We leave 5% room for error.
	assert.Less(t, math.Abs(results["Noam"]-0.333), 0.05)
	assert.Less(t, math.Abs(results["Bean"]-0.333), 0.05)
	assert.Less(t, math.Abs(results["TheGamblr"]-0.333), 0.05)
}
