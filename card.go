package TheGamblr

type Card struct {
	rank Rank
	suit Suit
}

func NewCard(rank Rank, suit Suit) *Card {
	return &Card{rank: rank, suit: suit}
}

func (c *Card) Name() string {
	return c.rank.Name() + c.suit.Name()
}

func StrToCard(name string) *Card {
	return &Card{
		rank: StrToRank(name[:1]),
		suit: StrToSuit(name[1:2]),
	}
}
