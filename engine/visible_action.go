package engine

type visibleAction struct {
	actionTaken Action
	player      ActivePlayerState
	amount      int
}
type VisibleAction interface {
	// ActionTaken represents which action was taken play Player
	ActionTaken() Action
	// Player is which player took this action
	Player() ActivePlayerState
	// Amount will be the amount of chips the Action refers too.
	Amount() int
}

func (v *visibleAction) ActionTaken() Action {
	return v.actionTaken
}

func (v *visibleAction) Player() ActivePlayerState {
	return v.player
}

func (v *visibleAction) Amount() int {
	return v.amount
}

func newVisibleAction(player ActivePlayerState, actionTaken Action, amount int) VisibleAction {
	return &visibleAction{
		actionTaken: actionTaken,
		player:      player,
		amount:      amount,
	}
}
