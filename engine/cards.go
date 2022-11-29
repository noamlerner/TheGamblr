package engine

type Cards []*Card

func (c Cards) Len() int      { return len(c) }
func (c Cards) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Cards) Less(i, j int) bool {
	if c[i].rank != c[j].rank {
		return c[i].rank > c[j].rank
	}
	return c[i].suit < c[j].suit
}

func StrToCards(cardsStr []string) Cards {
	cards := make([]*Card, len(cardsStr))
	for i, c := range cardsStr {
		cards[i] = StrToCard(c)
	}
	return cards
}
