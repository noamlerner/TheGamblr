package pokerengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	for s := 0; s < NumSuits; s++ {
		for r := 0; r < NumRanks; r++ {
			assert.Equal(t, deck.NextCard(), NewCard(IToRank(r), IToSuit(s)))
		}
	}
}

func TestShuffle(t *testing.T) {
	shuffledDeck := NewDeck().Shuffle()
	deck := NewDeck()
	// 1/52 chance of being flaky but i'm feeling lazy
	assert.NotEqual(t, deck.NextCard(), shuffledDeck.NextCard())
}
