package engine

type OneActionBot struct {
	action     ActionType
	numCalled  int
	boardState BoardState
}

func (c *OneActionBot) ReceiveCards(hand Cards, blind int, boardState BoardState) {
	c.boardState = boardState
}

func (c *OneActionBot) SeeBoardState(boardState BoardState) {
}

func (c *OneActionBot) Act() (ActionType, int) {
	c.numCalled++
	return c.action, 10
}

func (c *OneActionBot) ActionUpdate(action Action) {
}
