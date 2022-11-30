package engine

// playerForWinnerCalculations is used to calculate winners and who is going to show their hands.
type playerForWinnerCalculations struct {
	p    *playerState
	hand *Hand
	// true if the player is going to show their hand.
	willShowHand bool
	chipsWon     int
}

func (p *playerForWinnerCalculations) winChips(chips int) {
	p.p.winChips(chips)
	p.chipsWon += chips
}

func (p *playerForWinnerCalculations) toState() *roundEndStats {
	var cards Cards
	handStrength := HandStrengthUnset
	if p.willShowHand {
		cards = p.p.cards
		handStrength = p.hand.strength
	}
	return &roundEndStats{
		chipsWon:     p.chipsWon,
		cards:        cards,
		handStrength: handStrength,
	}
}

type EndRoundPlayers []*playerForWinnerCalculations

func (h EndRoundPlayers) Len() int      { return len(h) }
func (h EndRoundPlayers) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h EndRoundPlayers) Less(i, j int) bool {
	return h[i].hand.Beats(h[j].hand)
}
