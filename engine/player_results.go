package engine

// playerForWinnerCalculations is used to calculate winners and who is going to show their hands.
type playerForWinnerCalculations struct {
	p    *playerState
	hand *Hand
	// true if the player is going to show their hand.
	willShowHand bool
	chipsWon     int
}

// roundEndPlayerState will be seen by everyone. Cards and HandStrength may be left unset if the player mucked.
type roundEndPlayerState struct {
	activePlayerState
	chipsWon     int
	cards        Cards
	handStrength HandStrength
}

func (r *roundEndPlayerState) ChipsWon() int {
	return r.chipsWon
}

func (r *roundEndPlayerState) Cards() Cards {
	return r.cards
}

func (r *roundEndPlayerState) HandStrength() HandStrength {
	return r.handStrength
}

type PlayerResults interface {
	// Stack returns how many chips this player has to bet.
	Stack() int
	// Status returns one of the possible PlayerStatus
	Status() PlayerStatus
	// SeatNumber return the index of this player on the board
	SeatNumber() int
	// Id returns a unique player ID.
	Id() string
	// ChipsWon returns how many chips this player won, 0 if none.
	ChipsWon() int
	// Cards will by nil if the player mucked, otherwise you will see their hand.
	Cards() Cards
	// HandStrength will return the engine's calcualtion of their hand strength.
	HandStrength() HandStrength
}

func (p *playerForWinnerCalculations) winChips(chips int) {
	p.p.winChips(chips)
	p.chipsWon += chips
}

func (p *playerForWinnerCalculations) toState() PlayerResults {
	var cards Cards
	handStrength := HandStrengthUnset
	if p.willShowHand {
		cards = p.p.cards
		handStrength = p.hand.strength
	}
	return &roundEndPlayerState{
		activePlayerState: p.p.activePlayerState,
		chipsWon:          p.chipsWon,
		cards:             cards,
		handStrength:      handStrength,
	}
}

type EndRoundPlayers []*playerForWinnerCalculations

func (h EndRoundPlayers) Len() int      { return len(h) }
func (h EndRoundPlayers) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h EndRoundPlayers) Less(i, j int) bool {
	return h[i].hand.Beats(h[j].hand)
}
