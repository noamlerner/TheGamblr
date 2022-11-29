package TheGamblr

type Action int

const (
	FoldAction Action = iota
	CallAction
	RaiseAction
	CheckFold
	SmallBlind
	BigBlind
)

type BotPlayer interface {
	// ReceiveCards is a way to give every player the information as to what cards they will be playing with this
	// round.
	ReceiveCards(hand Cards, blind int, boardState ActiveBoard)
	// SeeActiveBoardState is called before players take action on PreFlop, Flop, Turn and River. It is also called once a
	// game has ended.
	SeeActiveBoardState(boardState ActiveBoard)
	// RoundResults shows the final board state with how many chips people won and shows peoples hands if possible.
	RoundResults(results RoundResults)
	// Act allows the player to return what action they want to take. The second return value is only considered
	// if Action == RaiseAction, in which case it is the amount to raise by. If the amount is more than is available for
	// the player, it will be considered AllIn. If the amount is less than the MinBet, we will automatically raise it to
	// the MinBet.
	Act() (Action, int)
	// ReceiveUpdate lets the bot know of an action another bot player took.
	ReceiveUpdate(action VisibleAction)
}

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
