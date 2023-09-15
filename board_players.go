package pokerengine

type BoardPlayers []*BoardPlayer
type BoardPlayerFunc func(p *BoardPlayer)

func (p BoardPlayers) IterateActive(first int, f BoardPlayerFunc) {
	p.iterateActive(first, len(p), f)
	p.iterateActive(0, first, f)
}

// IterateActiveUntil will start iterating from the first and iterate around the table until we reach
// last Exclusive. If dealerButton == endAt, that is the same as called IterateActive - we will just go all
// the way around.
func (p BoardPlayers) IterateActiveUntil(first int, last int, f BoardPlayerFunc) {
	if first == last {
		p.IterateActive(first, f)
	} else if last > first {
		p.iterateActive(first, last, f)
	} else {
		p.iterateActive(first, len(p), f)
		p.iterateActive(0, last, f)
	}
}

func (p BoardPlayers) iterateActive(first int, last int, f BoardPlayerFunc) {
	for i := first; i < last; i++ {
		if p[i] != nil && p[i].Status() != BoardPlayerStatusOut {
			f(p[i])
		}
	}
}
