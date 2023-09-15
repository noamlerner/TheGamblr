package engine

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddsOfWinning(t *testing.T) {
	tests := []struct {
		name           string
		hands          []Cards
		numPlayers     int
		communityCards Cards
		odds           []float64
	}{
		{
			name: "AK suited v JJ ",
			hands: []Cards{
				{NewCard(Ace, Diamonds), NewCard(King, Diamonds)},
				{NewCard(Jack, Hearts), NewCard(Jack, Clubs)},
			},
			numPlayers: 2,
			odds:       []float64{0.4594, 0.5368},
		},
		{
			name: "2 defined in 3 player",
			hands: []Cards{
				{NewCard(Eight, Hearts), NewCard(Queen, Hearts)},
				{NewCard(Ten, Diamonds), NewCard(Jack, Diamonds)},
			},
			numPlayers: 3,
			odds:       []float64{0.378, 0.379},
		},
		{
			name: "2 defined in 3 player with community card",
			hands: []Cards{
				{NewCard(Eight, Hearts), NewCard(Queen, Hearts)},
				{NewCard(Ten, Diamonds), NewCard(Jack, Diamonds)},
			},
			communityCards: Cards{NewCard(Jack, Spades)},
			numPlayers:     3,
			odds:           []float64{0.216, 0.632},
		},
		{
			name: "AA vs QQ 5 player",
			hands: []Cards{
				{NewCard(Ace, Diamonds), NewCard(Ace, Clubs)},
				{NewCard(Queen, Hearts), NewCard(Queen, Clubs)},
			},
			numPlayers: 5,
			odds:       []float64{0.534, 0.158},
		},
		{
			name: "4 players, 3 defined, 1 community card",
			hands: []Cards{
				{NewCard(Eight, Hearts), NewCard(King, Diamonds)},
				{NewCard(Seven, Clubs), NewCard(Six, Diamonds)},
				{NewCard(Four, Spades), NewCard(Five, Spades)},
			},
			communityCards: Cards{NewCard(Six, Spades)},
			numPlayers:     4,
			odds:           []float64{0.176, 0.351, 0.326},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			odds := OddsOfWinningOrTie(test.hands, test.numPlayers, test.communityCards, 100000)
			for i := range test.odds {
				assert.True(t, math.Abs(test.odds[i]-odds[i]) < 0.01, test.odds[i])
			}
		})
	}
}
