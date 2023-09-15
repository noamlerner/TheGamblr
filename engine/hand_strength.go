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
