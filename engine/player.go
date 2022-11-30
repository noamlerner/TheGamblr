package engine

type BotPlayer interface {
	// ReceiveCards is a way to give every player the information as to what cards they will be playing with this
	// round.
	ReceiveCards(hand Cards, blind int, boardState BoardState)
	// SeeBoardState is called before players take action on PreFlop, Flop, Turn and River. It is also called once a
	// game has ended.
	SeeBoardState(boardState BoardState)
	// Act allows the player to return what action they want to take. The second return value is only considered
	// if ActionType == RaiseAction, in which case it is the amount to raise by. If the amount is more than is available for
	// the player, it will be considered AllIn. If the amount is less than the MinBet, we will automatically raise it to
	// the MinBet.
	Act() (ActionType, int)
	// ActionUpdate lets the bot know of an action another bot player took.
	ActionUpdate(action Action)
}
