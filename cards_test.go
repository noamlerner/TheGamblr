package TheGamblr

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	cards := Cards([]*Card{
		NewCard(IToRank(1), IToSuit(1)),
		NewCard(IToRank(6), IToSuit(2)),
		NewCard(IToRank(4), IToSuit(0)),
		NewCard(IToRank(3), IToSuit(3)),
		NewCard(IToRank(12), IToSuit(0)),
	})

	sort.Sort(cards)

	assert.Equal(t, NewCard(IToRank(12), IToSuit(0)), cards[0])
	assert.Equal(t, NewCard(IToRank(6), IToSuit(2)), cards[1])
	assert.Equal(t, NewCard(IToRank(4), IToSuit(0)), cards[2])
	assert.Equal(t, NewCard(IToRank(3), IToSuit(3)), cards[3])
	assert.Equal(t, NewCard(IToRank(1), IToSuit(1)), cards[4])
}
