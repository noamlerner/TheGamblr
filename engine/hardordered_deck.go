package engine

// HardOrderedDeck is broken into 3 parts.
// 1. The ordering of cards we deal to players
// 2. The community cards
// 3. everything else
// This allows us to simulate decks where certain cards are always dealt and the rest are random, for example
// If you wanted to simulate a deck where the first player gets A, A and the flop is Q, Q, Q with 2 other players you
// would create this deck using
// NewHardOrderedDeck([]Cards{AA}, 3, []Cards{QQQ})
type HardOrderedDeck struct {
	deck

	playerHands Cards
	// Allows for other players without hardcoded hands to be in the game
	numFillerCards int
	communityCards Cards
	restOfCards    Cards

	// so we don't have to keep reallocating it
	fullDeck Cards
}

func NewHardOrderedDeck(hardCodedHands []Cards, totalNumPlayers int, hardCodedCommunityCards Cards) Deck {
	playerHands := make(Cards, 0, len(hardCodedHands)*2)
	cardsInOrdering := map[int]bool{}

	for _, hand := range hardCodedHands {
		for _, card := range hand {
			cardsInOrdering[card.ID()] = true
			playerHands = append(playerHands, card)
		}
	}

	for _, card := range hardCodedCommunityCards {
		cardsInOrdering[card.ID()] = true
	}

	restOfTheCards := make(Cards, 0, NumCardsInDeck-len(cardsInOrdering))
	for s := 0; s < NumSuits; s++ {
		for r := 0; r < NumRanks; r++ {
			card := NewCard(IToRank(r), IToSuit(s))
			_, ok := cardsInOrdering[card.ID()]
			if !ok {
				restOfTheCards = append(restOfTheCards, card)
			}
		}
	}

	d := NewDeck().(*deck)
	return &HardOrderedDeck{
		deck:           *d,
		playerHands:    playerHands,
		numFillerCards: (totalNumPlayers - len(hardCodedHands)) * 2,
		communityCards: hardCodedCommunityCards,
		fullDeck:       d.cards,
		restOfCards:    restOfTheCards,
	}
}

func (d *HardOrderedDeck) Shuffle() Deck {
	d.onCard = 0
	d.deck.cards = d.restOfCards
	d.deck.Shuffle()

	i := len(d.playerHands)
	copy(d.fullDeck[:i], d.playerHands)
	j := i + d.numFillerCards
	copy(d.fullDeck[i:j], d.deck.cards[:d.numFillerCards])
	i = j
	j = i + len(d.communityCards)
	copy(d.fullDeck[i:j], d.communityCards)
	copy(d.fullDeck[j:], d.deck.cards[d.numFillerCards:])
	d.deck.cards = d.fullDeck
	return d
}
