package engine

type OneActionBot struct {
	action     Action
	numCalled  int
	boardState ActiveBoard
}

func (c *OneActionBot) RoundResults(results RoundResults) {
}

func (c *OneActionBot) ReceiveCards(hand Cards, blind int, boardState ActiveBoard) {
	c.boardState = boardState
}

func (c *OneActionBot) SeeActiveBoardState(boardState ActiveBoard) {
}

func (c *OneActionBot) Act() (Action, int) {
	c.numCalled++
	return c.action, 10
}

func (c *OneActionBot) ReceiveUpdate(action VisibleAction) {
}
