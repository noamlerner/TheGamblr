package engine

type HandStrength int

const (
	HandStrengthUnset HandStrength = iota
	HighCard
	Pair
	TwoPair
	Trips
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
)

var (
	name = map[HandStrength]string{
		HighCard:      "HighCard",
		Pair:          "Pair",
		TwoPair:       "TwoPair",
		Trips:         "ThreeOfAKind",
		Straight:      "Straight",
		Flush:         "Flush",
		FullHouse:     "FullHouse",
		Quads:         "FourOfAKind",
		StraightFlush: "StraightFlush",
	}
)

func (h HandStrength) String() string {
	return name[h]
}
