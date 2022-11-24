package pokerengine

type HandStrength int

const (
	HighCard HandStrength = iota
	Pair
	TwoPair
	Trips
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
)
