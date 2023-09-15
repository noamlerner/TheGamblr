package pokerengine

type HandedBoardPlayer struct {
	p *BoardPlayer
	h *Hand
}

type HandedBoardPlayers []*HandedBoardPlayer

func (h HandedBoardPlayers) Len() int      { return len(h) }
func (h HandedBoardPlayers) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h HandedBoardPlayers) Less(i, j int) bool {
	return h[i].h.Beats(h[j].h)
}
