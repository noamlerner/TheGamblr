package engine

type Stage int

const NumStages = 4
const (
	PreFlop Stage = iota
	Flop
	Turn
	River
)

func (s Stage) nextStage() Stage {
	n := s + 1
	if n == NumStages {
		n = 0
	}
	return n
}
