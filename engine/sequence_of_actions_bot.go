package engine

type SequenceOfActionsBot struct {
	actions   []ActionType
	i         int
	numCalled int
}

func (c *SequenceOfActionsBot) ReceiveCards(hand Cards, blind int, boardState BoardState) {
}

func (c *SequenceOfActionsBot) SeeBoardState(boardState BoardState) {
}

func (c *SequenceOfActionsBot) Act() (ActionType, int) {
	c.numCalled++
	i := c.i
	c.i++
	if i >= len(c.actions) {
		return CallAction, 10
	}
	return c.actions[i], 10
}

func (c *SequenceOfActionsBot) ActionUpdate(action Action) {
}
