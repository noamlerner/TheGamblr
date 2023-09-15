package engine

type PlayerRoundResults interface {
	// ChipsWon returns how many chips this player won, 0 if none.
	ChipsWon() int
	// Cards will by nil if the player mucked, otherwise you will see their hand.
	Cards() Cards
	// HandStrength will return the engine's calcualtion of their hand strength.
	HandStrength() HandStrength
}
type PlayerState interface {
	// Stack returns how many chips this player has to bet.
	Stack() int
	// Status returns one of the possible PlayerStatus
	Status() PlayerStatus
	// SeatNumber return the index of this player on the board
	SeatNumber() int
	// Id returns a unique player ID.
	Id() string
	// PlayerRoundResults will be nil if this isnt a round end or if the player didn't make it there.
	RoundEndStats() PlayerRoundResults
}

// visiblePlayerState is the struct that is visible to everyone
type visiblePlayerState struct {
	stack         int
	status        PlayerStatus
	seatNumber    int
	id            string
	roundEndStats PlayerRoundResults
}

func (v *visiblePlayerState) Stack() int {
	return v.stack
}

func (v *visiblePlayerState) Status() PlayerStatus {
	return v.status
}

func (v *visiblePlayerState) SeatNumber() int {
	return v.seatNumber
}

func (v *visiblePlayerState) Id() string {
	return v.id
}

func (v *visiblePlayerState) RoundEndStats() PlayerRoundResults {
	return v.roundEndStats
}

type roundEndStats struct {
	chipsWon     int
	cards        Cards
	handStrength HandStrength
}

func (r *roundEndStats) ChipsWon() int {
	if r == nil {
		return 0
	}
	return r.chipsWon
}

func (r *roundEndStats) Cards() Cards {
	if r == nil {
		return Cards{}
	}
	return r.cards
}

func (r *roundEndStats) HandStrength() HandStrength {
	if r == nil {
		return HandStrengthUnset
	}
	return r.handStrength
}
