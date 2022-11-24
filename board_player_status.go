package pokerengine

type BoardPlayerStatus int

const (
	// BoardPlayerStatusOut means the player has lost all their money
	BoardPlayerStatusOut BoardPlayerStatus = iota
	// BoardPlayerStatusFolded means the player has folded the round
	BoardPlayerStatusFolded
	// BoardPlayerStatusPlaying means the player is in the round
	BoardPlayerStatusPlaying
	// BoardPlayerStatusAllIn is for players that are in the game, but are currently AllIn
	BoardPlayerStatusAllIn
)

func (s BoardPlayerStatus) EligibleToWin() bool {
	return s >= BoardPlayerStatusPlaying
}
