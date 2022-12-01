package engine

import (
	"sort"
)

// Hand is a set of 5 cards and a defining HandStrength. The first card in the cards slice is guaranteed to be the
// HighCard, but not the highest card. This distinction is only made in a straight to 5 where A is the highest card,
// but the HighCard is 5, so 5 will be the first card.
type Hand struct {
	cards    []*Card
	strength HandStrength

	boardPlayer *playerState
}

func NewHandForPlayer(board *boardState, boardPlayer *playerState) *Hand {
	hand := NewHand(append(board.communityCards, boardPlayer.cards...))
	hand.boardPlayer = boardPlayer
	return hand
}

func NewHand(allCards Cards) *Hand {
	sort.Sort(allCards)

	cardsByRank := map[Rank]Cards{}
	cardsBySuit := map[Suit]Cards{}
	for _, c := range allCards {
		cardsByRank[c.rank] = append(cardsByRank[c.rank], c)
		cardsBySuit[c.suit] = append(cardsBySuit[c.suit], c)
	}

	hasFlush := false
	var flushSuit Suit
	for suit, cards := range cardsBySuit {
		if len(cards) >= 5 {
			hasFlush = true
			flushSuit = suit
		}
	}

	// check for straight flush
	if hasFlush {
		straightFlush := extractStraight(cardsBySuit[flushSuit])
		if straightFlush != nil {
			return &Hand{
				cards:    straightFlush,
				strength: StraightFlush,
			}
		}
	}

	// Hands made up of repeated cards (Either Quads, Trips, Pairs, HighCard)
	repeatedRankHands := make(Hands, 0, len(cardsByRank))
	for _, cards := range cardsByRank {
		repeatedRankHands = append(repeatedRankHands, repeatedHandRank(cards))
	}
	sort.Sort(repeatedRankHands)

	if repeatedRankHands[0].strength == Quads {
		repeatedRankHands[0].fillInHighCards(allCards)
		return repeatedRankHands[0]
	}

	if repeatedRankHands[0].strength == Trips && len(repeatedRankHands) > 1 && repeatedRankHands[1].strength >= Pair {
		// Full house
		repeatedRankHands[0].strength = FullHouse
		repeatedRankHands[0].cards = append(repeatedRankHands[0].cards, repeatedRankHands[1].cards[:2]...)
		return repeatedRankHands[0]
	}

	if hasFlush {
		return &Hand{
			strength: Flush,
			cards:    cardsBySuit[flushSuit][:5],
		}
	}

	if straight := extractStraight(allCards); straight != nil {
		return &Hand{
			strength: Straight,
			cards:    straight,
		}
	}

	if repeatedRankHands[0].strength == Trips {
		repeatedRankHands[0].fillInHighCards(allCards)
		return repeatedRankHands[0]
	}

	if repeatedRankHands[0].strength == Pair && repeatedRankHands.Len() > 1 && repeatedRankHands[1].strength == Pair {
		repeatedRankHands[0].strength = TwoPair
		repeatedRankHands[0].cards = append(repeatedRankHands[0].cards, repeatedRankHands[1].cards...)
		repeatedRankHands[0].fillInHighCards(allCards)
		return repeatedRankHands[0]
	}

	repeatedRankHands[0].fillInHighCards(allCards)

	return repeatedRankHands[0]
}

func (h *Hand) fillInHighCards(cards Cards) {
	ranks := map[string]bool{}
	for _, c := range h.cards {
		ranks[c.Name()] = true
	}
	for _, c := range cards {
		_, ok := ranks[c.Name()]
		if !ok {
			h.cards = append(h.cards, c)
		}

		if len(h.cards) == 5 {
			return
		}
	}
}

// repeatedHandRank returns a Hand which is made up of all the passed in cards.
// The passed in cards are assumed to be all the same rank. We will define the strength based on the length
// of the slice. Len(4) == Quads... etc
func repeatedHandRank(cards []*Card) *Hand {
	handStrength := HighCard
	switch len(cards) {
	case 4:
		handStrength = Quads
	case 3:
		handStrength = Trips
	case 2:
		handStrength = Pair
	}
	return &Hand{
		strength: handStrength,
		cards:    cards,
	}
}

// Beats returns true if hand is a stronger hand than o
func (h *Hand) Beats(o *Hand) bool {
	if h.strength > o.strength {
		return true
	}
	if o.strength != h.strength {
		return false
	}

	for i, c := range h.cards {
		if c.rank > o.cards[i].rank {
			return true
		}
		if c.rank < o.cards[i].rank {
			return false
		}
	}
	return false
}

// Tie returns true if hand is the same strength as o
func (h *Hand) Tie(o *Hand) bool {
	if o.strength != h.strength {
		return false
	}

	for i, c := range h.cards {
		if c.rank != o.cards[i].rank {
			return false
		}
	}
	return true
}

// extractStraight returns the best 5 cards for a straight. If they do not exist, this will return nil.
func extractStraight(cards Cards) []*Card {
	if len(cards) < 5 {
		return nil
	}
	chainLength := 1
	for i := 1; i < len(cards); i++ {
		if cards[i].rank+1 == cards[i-1].rank {
			chainLength++

			if chainLength == 5 {
				return cards[i-4 : i+1]
			}

			continue
		}

		if cards[i].rank == cards[i-1].rank {
			cards = append(cards[:i], cards[i+1:]...)
			i--
			continue
		}

		chainLength = 1
	}

	if chainLength == 4 && cards[len(cards)-1].rank == Two && cards[0].rank == Ace {
		// A->5
		return append(cards[len(cards)-4:], cards[0])
	}
	return nil
}
