package engine

type SequenceOfActionsBot struct {
	sequence  []ActionType
	i         int
	numCalled int
}

func (c *SequenceOfActionsBot) ReceiveCards(hand Cards) {
}

func (c *SequenceOfActionsBot) SeeBoardState(boardState BoardState) {
}

func (c *SequenceOfActionsBot) Act(int, int, int) (ActionType, int) {
	c.numCalled++
	i := c.i
	c.i++
	if i >= len(c.sequence) {
		return CallAction, 10
	}
	return c.sequence[i], 10
}

func (c *SequenceOfActionsBot) ActionUpdate(action Action, state BoardState) {
}
