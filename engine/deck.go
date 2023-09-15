package engine

import (
	"math/rand"
	"time"
)

const NumCardsInDeck = 52

type Deck interface {
	NextCard() *Card
	NextCards(n int) []*Card
	Shuffle() Deck
	Cards() Cards
}

type deck struct {
	cards  []*Card
	onCard int
	rand   *rand.Rand

	// if these are set, we will place the cards at the front of the deck. good for simulations.
	cardsToTheFront Cards
	restOfTheDeck   Cards
}

func (d *deck) Cards() Cards {
	return d.cards
}

func NewDeckWithCards(cards []*Card) Deck {
	return &deck{cards: cards, rand: rand.New(rand.NewSource(time.Now().Unix()))}

}
func NewDeck() Deck {
	cards := make([]*Card, NumCardsInDeck)
	for s := 0; s < NumSuits; s++ {
		for r := 0; r < NumRanks; r++ {
			i := s*NumRanks + r
			cards[i] = NewCard(IToRank(r), IToSuit(s))
		}
	}
	return NewDeckWithCards(cards)
}

func (d *deck) NextCard() *Card {
	if d.onCard > NumCardsInDeck {
		return nil
	}
	c := d.cards[d.onCard]
	d.onCard++
	return c
}

func (d *deck) NextCards(n int) []*Card {
	if d.onCard > NumCardsInDeck {
		return nil
	}
	c := d.cards[d.onCard : d.onCard+n]
	d.onCard += n
	return c
}

func (d *deck) Shuffle() Deck {
	d.onCard = 0
	d.rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
	return d
}
