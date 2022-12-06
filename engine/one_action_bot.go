package engine

type OneActionBot struct {
	action     ActionType
	numCalled  int
	boardState BoardState
}

func (c *OneActionBot) ReceiveCards(hand Cards) {
}

func (c *OneActionBot) SeeBoardState(boardState BoardState) {
	c.boardState = boardState
}

func (c *OneActionBot) Act(int, int, int) (ActionType, int) {
	c.numCalled++
	return c.action, 10
}

func (c *OneActionBot) ActionUpdate(action Action, state BoardState) {
}
