package engine

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandAnalysis(t *testing.T) {
	NewHand([]*Card{NewCard(6, 2)})
}

func TestHasStraight(t *testing.T) {
	tests := []struct {
		name        string
		ranks       []Rank
		hasStraight bool
	}{
		{
			"Has Straight",
			[]Rank{0, 1, 2, 3, 4, 3, 8},
			true,
		},
		{
			"Does not have straight",
			[]Rank{0, 3, 5, 7, 6, 0, 0},
			false,
		},
		{
			"A to 5",
			[]Rank{Ace, Two, Three, Four, Five, Nine, Ten},
			true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cards := make([]*Card, len(test.ranks))
			for i, r := range test.ranks {
				cards[i] = NewCard(r, IToSuit(i%NumSuits))
			}
			sort.Sort(Cards(cards))

			assert.Equal(t, test.hasStraight, extractStraight(cards) != nil)
		})
	}
}

func TestBeats(t *testing.T) {
	tests := []struct {
		name  string
		h     *Hand
		o     *Hand
		beats bool
	}{
		{
			name: "HighCard vs StraightFlush",
			h: &Hand{
				cards:    GenerateHighCard(8),
				strength: HighCard,
			},
			o: &Hand{
				cards:    GenerateHighCard(7),
				strength: StraightFlush,
			},
			beats: false,
		},
		{
			name: "StraightFlush vs HighCard",
			o: &Hand{
				cards:    GenerateHighCard(8),
				strength: HighCard,
			},
			h: &Hand{
				cards:    GenerateHighCard(7),
				strength: StraightFlush,
			},
			beats: true,
		},
		{
			name: "Flush vs StraightFlush",
			h: &Hand{
				cards:    GenerateFlush(8),
				strength: Flush,
			},
			o: &Hand{
				cards:    GenerateHighCard(7),
				strength: StraightFlush,
			},
			beats: false,
		},
		{
			name: "The Higher Straight Flush",
			h: &Hand{
				cards:    GenerateFlush(8),
				strength: StraightFlush,
			},
			o: &Hand{
				cards:    GenerateHighCard(7),
				strength: StraightFlush,
			},
			beats: true,
		},
		{
			name: "The Lower Straight Flush",
			h: &Hand{
				cards:    GenerateFlush(6),
				strength: StraightFlush,
			},
			o: &Hand{
				cards:    GenerateHighCard(7),
				strength: StraightFlush,
			},
			beats: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.beats, test.h.Beats(test.o))
		})
	}
}

func TestNewHand(t *testing.T) {
	tests := []struct {
		name         string
		allCards     []string
		handStrength HandStrength
		resCards     []string
		rankOnly     bool
	}{
		{
			name:         "StraightFlush",
			allCards:     []string{"AH", "2H", "6D", "8S", "3H", "4H", "5H"},
			handStrength: StraightFlush,
			resCards:     []string{"5H", "4H", "3H", "2H", "AH"},
		},
		{
			name:         "Quads",
			allCards:     []string{"6H", "6D", "AS", "6S", "5D", "6C", "4S"},
			handStrength: Quads,
			resCards:     []string{"6H", "6S", "6D", "6C", "AS"},
		},
		{
			name:         "FullHouse",
			allCards:     []string{"2H", "6S", "2S", "6H", "3H", "6D", "7H"},
			handStrength: FullHouse,
			resCards:     []string{"6S", "2H", "2S", "6H", "6D"},
		},
		{
			name:         "Flush",
			allCards:     []string{"JH", "QH", "2H", "6H", "AH", "3S", "7H"},
			handStrength: Flush,
			resCards:     []string{"AH", "JH", "QH", "7H", "6H"},
		},
		{
			name:         "Straight",
			allCards:     []string{"KH", "QD", "JC", "AS", "TH", "9D", "6C"},
			handStrength: Straight,
			resCards:     []string{"AS", "KH", "QD", "JC", "TH"},
		},
		{
			name:         "Trips",
			allCards:     []string{"6H", "6D", "6S", "9C", "8H", "JD", "QS"},
			handStrength: Trips,
			resCards:     []string{"6H", "6D", "6S", "JD", "QS"},
		},
		{
			name:         "TwoPair",
			allCards:     []string{"5H", "5C", "2H", "2C", "JD", "TD", "8D"},
			handStrength: TwoPair,
			resCards:     []string{"5H", "5C", "2H", "2C", "JD"},
		},
		{
			name:         "Pair",
			allCards:     []string{"5H", "5C", "AH", "2C", "JD", "TD", "8D"},
			handStrength: Pair,
			resCards:     []string{"5H", "5C", "AH", "TD", "JD"},
		},
		{
			name:         "HighCard",
			allCards:     []string{"KH", "QC", "8D", "7S", "3S", "2C", "5D"},
			handStrength: HighCard,
			resCards:     []string{"KH", "QC", "8D", "7S", "5D"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			allCards := NewDeckWithCards(StrToCards(test.allCards)).Shuffle().Cards()
			hand := NewHand(allCards)
			assert.Equal(t, test.handStrength, hand.strength)
			resCards := hand.cards
			expectedCards := StrToCards(test.resCards)
			assert.Equal(t, expectedCards[0].rank, resCards[0].rank)
			assert.Equal(t, len(expectedCards), len(resCards))

			cards := map[string]bool{}
			for _, c := range expectedCards {
				cards[c.Name()] = true
			}

			for _, c := range resCards {
				assert.Contains(t, cards, c.Name())
			}
		})
	}
}
