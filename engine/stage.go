package engine

type Stage int

const NumStages = 4
const (
	PreFlop Stage = iota
	Flop
	Turn
	River
	RoundOver
)

var (
	stageName = map[Stage]string{
		PreFlop:   "PreFlop",
		Flop:      "Flop",
		Turn:      "Turn",
		River:     "River",
		RoundOver: "RoundOver",
	}
)

func (s Stage) nextStage() Stage {
	n := s + 1
	if n == NumStages {
		n = 0
	}
	return n
}

func (s Stage) Name() string {
	return stageName[s]
}
