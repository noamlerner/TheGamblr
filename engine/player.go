package engine

type BotPlayer interface {
	// ReceiveCards is a way to give every player the information as to what cards they will be playing with this
	// round.
	ReceiveCards(hand Cards, blind int)
	// SeeBoardState is called before players take action on PreFlop, Flop, Turn and River. It is also called once a
	// game has ended.
	SeeBoardState(boardState BoardState)
	// Act allows the player to return what action they want to take. The second return value is only considered
	// if ActionType == RaiseAction, in which case it is the amount to raise by. If the amount is more than is available for
	// the player, it will be considered AllIn. If the amount is less than the MinBet, we will automatically raise it to
	// the MinBet.
	// Three ints are provided as input.
	// One: The current pot with everyone's bet put in.
	// Two: The total call amount (i.e. every player has to put in 100 chips).
	// Three: The amount left for this particular player to call (you already put in 50, this would be 50).
	Act(pot int, callAmount int, leftToCall int) (ActionType, int)
	// ActionUpdate lets the bot know of an action another bot player took.
	ActionUpdate(action Action)
}
