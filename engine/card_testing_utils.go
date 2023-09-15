package engine

func GenerateRandHand() []*Card {
	return NewDeck().Shuffle().NextCards(5)
}

func GenerateStraightTo(high Rank) []*Card {
	cards := make([]*Card, 5)
	cards[0] = NewCard(high, IToSuit(0))
	for i := 1; i < 5; i++ {
		cards[i] = NewCard(cards[i-1].rank.PrevRank(), IToSuit(i%NumSuits))
	}
	return cards
}

func GenerateHighCard(high Rank) []*Card {
	cards := make([]*Card, 5)
	cards[0] = NewCard(high, IToSuit(0))
	cards[1] = NewCard(high.PrevRank().PrevRank(), IToSuit(1))
	for i := 2; i < 5; i++ {
		cards[i] = NewCard(cards[i-1].rank.PrevRank(), IToSuit(i%NumSuits))
	}
	return cards
}

func GenerateStraightFlush(high Rank) []*Card {
	cards := make([]*Card, 5)
	cards[0] = NewCard(high, IToSuit(0))
	for i := 1; i < 5; i++ {
		cards[i] = NewCard(cards[i-1].rank.PrevRank(), IToSuit(0))
	}
	return cards
}

func GenerateFlush(high Rank) []*Card {
	cards := make([]*Card, 5)
	cards[0] = NewCard(high, IToSuit(0))
	cards[1] = NewCard(high.PrevRank().PrevRank(), IToSuit(0))
	for i := 2; i < 5; i++ {
		cards[i] = NewCard(cards[i-1].rank.PrevRank(), IToSuit(0))
	}
	return cards
}
