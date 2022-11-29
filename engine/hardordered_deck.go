package engine

type HardOrderedDeck struct {
	deck
	ordering    Cards
	restOfCards Cards
}

func NewHardOrderedDeck(ordering Cards) Deck {
	cardsInFront := map[int]bool{}
	for _, card := range ordering {
		cardsInFront[card.ID()] = true
	}

	restOfTheCards := make(Cards, 0, NumCardsInDeck-len(ordering))
	for s := 0; s < NumSuits; s++ {
		for r := 0; r < NumRanks; r++ {
			card := NewCard(IToRank(r), IToSuit(s))
			if !cardsInFront[card.ID()] {
				restOfTheCards = append(restOfTheCards)
			}
		}
	}

	d := NewDeckWithCards(append(ordering, restOfTheCards...)).(*deck)

	return &HardOrderedDeck{deck: *d, ordering: ordering, restOfCards: restOfTheCards}
}

func (d *HardOrderedDeck) Shuffle() Deck {
	d.onCard = 0
	d.deck.cards = d.restOfCards
	d.deck.Shuffle()
	d.deck.cards = append(d.ordering, d.restOfCards...)
	return d
}
