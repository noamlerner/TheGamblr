package engine

import (
	"math/rand"
	"time"
)

const NumCardsInDeck = 52

type Deck struct {
	cards  []*Card
	onCard int
	rand   *rand.Rand
}

func NewDeckWithCards(cards []*Card) *Deck {
	return &Deck{cards: cards, rand: rand.New(rand.NewSource(time.Now().Unix()))}

}
func NewDeck() *Deck {
	cards := make([]*Card, NumCardsInDeck)
	for s := 0; s < NumSuits; s++ {
		for r := 0; r < NumRanks; r++ {
			i := s*NumRanks + r
			cards[i] = NewCard(IToRank(r), IToSuit(s))
		}
	}
	return NewDeckWithCards(cards)
}

func (d *Deck) NextCard() *Card {
	if d.onCard > NumCardsInDeck {
		return nil
	}
	c := d.cards[d.onCard]
	d.onCard++
	return c
}

func (d *Deck) NextCards(n int) []*Card {
	if d.onCard > NumCardsInDeck {
		return nil
	}
	c := d.cards[d.onCard : d.onCard+n]
	d.onCard += n
	return c
}

func (d *Deck) Shuffle() *Deck {
	d.onCard = 0
	d.rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
	return d
}
