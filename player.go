package pokerengine

type Action int

const (
	FoldAction Action = iota
	CallAction
	RaiseAction
	CheckFold
)

type BotPlayer interface {
	// ReceiveCards is a way to give every player the information as to what cards they will be playing with this
	// round.
	ReceiveCards(hand Cards, blind int)
	SeeCommunityCards(Stage, Cards)
	// Act allows the player to return what action they want to take. The second return value is only considered
	// if Action == RaiseAction, in which case it is the amount to raise by. If the amount is more than is available for
	// the player, it will be considered AllIn. If the amount is less than the MinBet, we will automatically raise it to
	// the MinBet.
	Act() (Action, int)
	ReceiveUpdate()
}

type OneActionBot struct {
	action    Action
	numCalled int
}

func (c *OneActionBot) ReceiveCards(hand Cards, blind int) {
}

func (c *OneActionBot) SeeCommunityCards(stage Stage, cards Cards) {
}

func (c *OneActionBot) Act() (Action, int) {
	c.numCalled++
	return c.action, 10
}

func (c *OneActionBot) ReceiveUpdate() {
}
