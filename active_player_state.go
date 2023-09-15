package TheGamblr

// activePlayerState is the struct that is visible to everyone
type activePlayerState struct {
	stack      int
	status     BoardPlayerStatus
	seatNumber int
	id         string
}

type ActivePlayerState interface {
	// Stack returns how many chips this player has to bet.
	Stack() int
	// Status returns one of the possible BoardPlayerStatus
	Status() BoardPlayerStatus
	// SeatNumber return the index of this player on the board
	SeatNumber() int
	// Id returns a unique player ID.
	Id() string
}

func (v *activePlayerState) Stack() int {
	return v.stack
}

func (v *activePlayerState) Status() BoardPlayerStatus {
	return v.status
}

func (v *activePlayerState) SeatNumber() int {
	return v.seatNumber
}

func (v *activePlayerState) Id() string {
	return v.id
}
