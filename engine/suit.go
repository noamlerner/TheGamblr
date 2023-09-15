package engine

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Spades
	Clubs
)
const NumSuits = 4

var (
	suitName = map[Suit]string{
		Hearts:   "Hearts",
		Diamonds: "Diamonds",
		Spades:   "Spades",
		Clubs:    "Clubs",
	}
	suitShortStr = map[Suit]string{
		Hearts:   "H",
		Diamonds: "D",
		Spades:   "S",
		Clubs:    "C",
	}
	reverseShortStr = map[string]Suit{
		"H": Hearts,
		"D": Diamonds,
		"S": Spades,
		"C": Clubs,
	}
)

// IToSuit lets you iterate over suits by converting i =0..3 to suits.
// 0 = Hearts
// 1 = Diamonds
// 2 = Spades
// 3 = Clubs
func IToSuit(i int) Suit {
	return Suit(i)
}
func (s Suit) Name() string {
	return suitName[s]
}

func (s Suit) ShortString() string {
	return suitShortStr[s]
}
func StrToSuit(shortStr string) Suit {
	return reverseShortStr[shortStr]
}
