package TheGamblr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceiveCards(t *testing.T) {
	tests := []struct {
		name        string
		blindAmount int
	}{
		{
			"Blind",
			10,
		},
		{
			"No Blind",
			0,
		},
		{
			"AllIn",
			200,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			player := &BoardPlayer{
				activePlayerState: activePlayerState{
					stack:  100,
					status: BoardPlayerStatusPlaying,
				},
				actor: &OneActionBot{action: CallAction},
			}

			cards := GenerateRandHand()[:2]
			blindPaid := player.ReceiveCards(cards, test.blindAmount, nil)
			assert.Equal(t, MinInt(test.blindAmount, 100), blindPaid)
			assert.Equal(t, MaxInt(0, 100-test.blindAmount), player.stack)
			assert.Equal(t, Cards(cards), player.cards)

			if player.stack == 0 {
				assert.Equal(t, BoardPlayerStatusAllIn, player.Status())
			} else {
				assert.Equal(t, BoardPlayerStatusPlaying, player.Status())
			}
		})
	}
}
