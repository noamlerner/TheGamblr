package engine

type SequenceOfActionsBot struct {
	actions []Action
	i       int
}

func (c *SequenceOfActionsBot) RoundResults(results RoundResults) {
}

func (c *SequenceOfActionsBot) ReceiveCards(hand Cards, blind int, boardState ActiveBoard) {
}

func (c *SequenceOfActionsBot) SeeActiveBoardState(boardState ActiveBoard) {
}

func (c *SequenceOfActionsBot) Act() (Action, int) {
	i := c.i
	c.i++
	if i >= len(c.actions) {
		return CallAction, 10
	}
	return c.actions[i], 10
}

func (c *SequenceOfActionsBot) ReceiveUpdate(action VisibleAction) {
}
