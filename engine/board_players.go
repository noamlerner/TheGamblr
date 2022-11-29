package engine

type PlayerStates []*playerState
type PlayerStateFunc func(p *playerState)

func (p PlayerStates) iterateActive(first int, f PlayerStateFunc) {
	p.iterateActive_(first, len(p), f)
	p.iterateActive_(0, first, f)
}

// iterateActiveUntil will start iterating from the first and iterate around the table until we reach
// last Exclusive. If dealerButton == endAt, that is the same as called iterateActive - we will just go all
// the way around.
func (p PlayerStates) iterateActiveUntil(first int, last int, f PlayerStateFunc) {
	if first == last {
		p.iterateActive(first, f)
	} else if last > first {
		p.iterateActive_(first, last, f)
	} else {
		p.iterateActive_(first, len(p), f)
		p.iterateActive_(0, last, f)
	}
}

func (p PlayerStates) iterateActive_(first int, last int, f PlayerStateFunc) {
	for i := first; i < last; i++ {
		if p[i] != nil && p[i].Status() != PlayerStatusOut {
			f(p[i])
		}
	}
}
