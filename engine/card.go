package engine

import (
	"fmt"
)

type Card struct {
	rank Rank
	suit Suit
}

func (c *Card) Rank() Rank {
	return c.rank
}

func (c *Card) Suit() Suit {
	return c.suit
}

const cardNameFormat = "%s of %s"

func NewCard(rank Rank, suit Suit) *Card {
	return &Card{rank: rank, suit: suit}
}

func (c *Card) Name() string {
	return fmt.Sprintf(cardNameFormat, c.rank.Name(), c.suit.Name())
}

func StrToCard(name string) *Card {
	return &Card{
		rank: StrToRank(name[:1]),
		suit: StrToSuit(name[1:2]),
	}
}

func (c *Card) ID() int {
	return int(c.suit)*NumRanks + int(c.rank)
}
