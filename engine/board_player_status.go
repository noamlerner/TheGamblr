package engine

type PlayerStatus int

const (
	// PlayerStatusOut means the player has lost all their money
	PlayerStatusOut PlayerStatus = iota
	// PlayerStatusFolded means the player has folded the round
	PlayerStatusFolded
	// PlayerStatusPlaying means the player is in the round
	PlayerStatusPlaying
	// PlayerStatusAllIn is for players that are in the game, but are currently AllIn
	PlayerStatusAllIn
)

func (s PlayerStatus) eligibleToWin() bool {
	return s >= PlayerStatusPlaying
}
