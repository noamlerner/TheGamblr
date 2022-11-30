package engine

type visibleAction struct {
	actionTaken ActionType
	player      PlayerState
	amount      int
}
type Action interface {
	// Type represents which action was taken play Player
	Type() ActionType
	// Player is which player took this action
	Player() PlayerState
	// Amount will be the amount of chips the ActionType refers too.
	Amount() int
}

func NewAction(t ActionType, p PlayerState, a int) Action {
	return &visibleAction{
		actionTaken: t,
		player:      p,
		amount:      a,
	}
}

func (v *visibleAction) Type() ActionType {
	return v.actionTaken
}

func (v *visibleAction) Player() PlayerState {
	return v.player
}

func (v *visibleAction) Amount() int {
	return v.amount
}

func newVisibleAction(player PlayerState, actionTaken ActionType, amount int) Action {
	return &visibleAction{
		actionTaken: actionTaken,
		player:      player,
		amount:      amount,
	}
}
