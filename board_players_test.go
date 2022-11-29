package TheGamblr

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardPlayers_IterateActive(t *testing.T) {
	tests := []struct {
		name           string
		firstPosition  int
		playersAtSeats []int
	}{
		{
			"Starting at 0",
			0,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			"Starting at 7",
			7,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			"Starting at 3, skip inactive",
			3,
			[]int{0, 1, 3, 4, 5, 7},
		},
		{
			"Starting at inactive",
			7,
			[]int{1, 2, 3, 4, 5},
		},
		{
			"one active",
			0,
			[]int{7},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			players, seen := setupPlayers(test.playersAtSeats)

			players.IterateActive(test.firstPosition, func(p *BoardPlayer) {
				seen[p.seatNumber] = true
			})

			for _, seat := range test.playersAtSeats {
				assert.True(t, seen[seat])
			}
		})
	}
}

func TestBoardPlayers_IterateActiveUntil(t *testing.T) {
	tests := []struct {
		name                        string
		firstPosition, lastPosition int
		playersAtSeats              []int
	}{
		{
			"0->7",
			0, 7,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			"7->0",
			7, 0,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			"6->1",
			6, 1,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			"5->5",
			5, 5,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			"5->4",
			5, 4,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},

		{
			"4->5",
			4, 5,
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
		},
		{
			"1->1 only one",
			1, 1,
			[]int{1},
		},
		{
			"3->6 Missing",
			3, 6,
			[]int{1, 2, 4, 7},
		},
		{
			"3->6 No PlayerResults",
			3, 6,
			[]int{1, 2, 7},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			players, seen := setupPlayers(test.playersAtSeats)

			players.IterateActiveUntil(test.firstPosition, test.lastPosition, func(p *BoardPlayer) {
				seen[p.seatNumber] = true
			})

			for _, seat := range test.playersAtSeats {
				if isBetween(test.firstPosition, test.lastPosition, seat) {
					assert.True(t, seen[seat])
				} else {
					assert.False(t, seen[seat])
				}
			}
		})
	}
}

func isBetween(first, last, check int) bool {
	if first < last {
		return check >= first && check < last
	} else if last < first {
		return check < last || check >= first
	} else {
		return true
	}
}

func setupPlayers(playersAtSeats []int) (BoardPlayers, map[int]bool) {
	players := make(BoardPlayers, 8)
	seen := map[int]bool{}
	for _, seat := range playersAtSeats {
		players[seat] = &BoardPlayer{
			activePlayerState: activePlayerState{
				id:         strconv.Itoa(seat),
				stack:      100,
				seatNumber: seat,
				status:     BoardPlayerStatusPlaying,
			},
			actor: &OneActionBot{action: CallAction},
		}
		seen[seat] = false
	}
	return players, seen
}
