package engine

import (
	"fmt"
)

type Cards []*Card

func (c Cards) Len() int      { return len(c) }
func (c Cards) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Cards) Less(i, j int) bool {
	if c[i].rank != c[j].rank {
		return c[i].rank > c[j].rank
	}
	return c[i].suit < c[j].suit
}

const (
	twoCardPlaceHolder   = "%s and %s"
	threeCardPlaceHolder = "%s, %s and %s"
	fourCardPlaceHolder  = "%s, %s, %s and %s"
	fiveCardPlaceHolder  = "%s, %s, %s, %s and %s"
)

func (c Cards) String() string {
	// we do it this way because there will normally be between 2 and 5 cards, so this
	// is a performance improvement.
	switch len(c) {
	case 2:
		return fmt.Sprintf(twoCardPlaceHolder, c[0].Name(), c[1].Name())
	case 3:
		return fmt.Sprintf(threeCardPlaceHolder, c[0].Name(), c[1].Name(), c[2].Name())
	case 4:
		return fmt.Sprintf(fourCardPlaceHolder, c[0].Name(), c[1].Name(), c[2].Name(), c[3].Name())
	case 5:
		return fmt.Sprintf(fiveCardPlaceHolder, c[0].Name(), c[1].Name(), c[2].Name(), c[3].Name(), c[4].Name())
	default:
		break
	}
	str := ""
	for _, card := range c {
		str += card.Name() + " "
	}
	return str
}

func StrToCards(cardsStr []string) Cards {
	cards := make([]*Card, len(cardsStr))
	for i, c := range cardsStr {
		cards[i] = StrToCard(c)
	}
	return cards
}
