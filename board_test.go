package pokerengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_NextActiveSeat(t *testing.T) {
	tests := []struct {
		name             string
		startingPosition int
		playersAtSeats   []int
	}{
		{
			"Starting at 0",
			0,
			[]int{1, 2, 3, 4, 5, 6, 7, 0},
		},
		{
			"Starting at 7",
			7,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			"Starting at 3, skip inactive",
			3,
			[]int{4, 5, 7, 0, 1, 3},
		},
		{
			"Starting at inactive",
			7,
			[]int{1, 2, 3, 4, 5},
		},
		{
			"one active - other",
			0,
			[]int{7},
		},

		{
			"one active - same ",
			0,
			[]int{0},
		},
		{
			"two active ",
			0,
			[]int{3, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)
			board := NewBoard()
			board.players = players
			pos := test.startingPosition
			for i := 0; i < len(test.playersAtSeats); i++ {
				assert.Equal(t, test.playersAtSeats[i], board.NextActiveSeat(pos))
				pos = test.playersAtSeats[i]
			}
		})
	}
}
