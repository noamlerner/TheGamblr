package engine

import "fmt"

type Rank int

const NumRanks = 13

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var (
	rankName = map[Rank]string{
		Two:   "2",
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
		Nine:  "9",
		Ten:   "T",
		Jack:  "J",
		Queen: "Q",
		King:  "K",
		Ace:   "A",
	}
	reverseRankName = map[string]Rank{
		"2": Two,
		"3": Three,
		"4": Four,
		"5": Five,
		"6": Six,
		"7": Seven,
		"8": Eight,
		"9": Nine,
		"T": Ten,
		"J": Jack,
		"Q": Queen,
		"K": King,
		"A": Ace,
	}
)

// IToRank returns a rank you can iterate over starting at 0. 0=2, 1=3...
func IToRank(i int) Rank {
	if i < 0 || i >= NumRanks {
		panic(fmt.Sprintf("I %v cannot be converted to rank", i))
	}
	return Rank(i)
}

func (r Rank) Name() string {
	return rankName[r]
}

func (r Rank) PrevRank() Rank {
	if r-1 >= 2 {
		return r - 1
	}
	return NumRanks - 1
}

func StrToRank(str string) Rank {
	return reverseRankName[str]
}
func (r Rank) NextRank() Rank {
	if r+1 < NumRanks {
		return r + 1
	}
	return Two
}
