package TheGamblr

type Stage int

const NumStages = 4
const (
	PreFlop Stage = iota
	Flop
	Turn
	River
)

func (s Stage) NextStage() Stage {
	n := s + 1
	if n == NumStages {
		n = 0
	}
	return n
}
