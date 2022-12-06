# The Gamblr

_Legend has it that this entire codebase was written in one playthrough
of [The Gambler by Kenny Rogers](https://www.youtube.com/watch?v=7hx4gdlfamo)_

...Those who know the author know he probably just put it on repeat for the whole time.

## Intro

_You gotta know when to hold 'em, know when to fold 'em_

This is a poker engine written to allow people to pit algorithms against each other. While this engine is written in Go,
it allows its users to use any language they want by creating a GRPC interface through which the game will be
played.

## How It Works

_If you're gonna play the game, boy, you gotta learn to play it right_

Poker playing algorithms can be written in any language. They will all interact with the `Casino`, a server which takes
grpc requests. It is recommended you separate out the logic for interacting with the `Casino` and the logic for actually
making your poker decisions.

### How to interact with the casino

_You never count your money when you're sittin' at the table_

To start the casino you can run `make casino` in the root directory (`make casino.kill` to kill it).

**If you want to see an example of a game, you can see a full working client in [client main](./client/main/main.go)
where we have two `ClientBot` wrapping around algorithms that take random actions play against each other.**

All interactions with the casino are defined in [casino.proto](./proto/casino.proto). The flow to start a game looks
like so:

1. One player needs to create a game using the CreateGame request, they receive the gameID
2. All players (including the creator) need to join a game using the gameID. The response to this request will include a
   gameToken that is required for every future request.
3. Once all players have joined, any player needs to start a game using the StartGame request.

Once a game has started, all players are expected to continually ask the casino for updates using the `ReceiveUpdates`
grpc. The response will provide a sequence number which you can put into your next request. This allows you to get only
the most recent updates.

#### ReceiveUpdates

The `ReceiveUpdates` response will contain an array of updates. Each of these can be one of two objects:

1. BoardState - provided at the beginning of each stage (PreFlop, Flop, Turn, River) and at the end of the game.
   Contains information about players in the hand, the pot and the community cards.
2. Action - Provided after everytime a player takes an action.

The final BoardState will contain info as to who won the round and what cards they had (unless they mucked their hand).

The ReceiveUpdates response will also contain your cards so you can make decisions.

Last it will contain the `MyActionPacket`. This will be null unless it is your turn. If it is your turn, this will
tell you the pot size, how much each player has to call, and how much you have left to call.
If you see a `MyActionPacket`, you are expected to call the `Act` grpc with the action you want to take, and chips you
want to put in. Note: If you try to Call/Raise, but you don't specify enough chips, the game will automatically raise
the amount for you.

### Writing Your Bot

_For a taste of your whiskey I'll give you some advice_

You can write your bot anyway you want as long as it fits into to the GRPC logic. The casino was written with a specific
interface in mind, and I believe it would be easiest for you to conform to this interface.

```
type BotPlayer interface {
	// ReceiveCards is a way to give every player the information as to what cards they will be playing with this
	// round.
	ReceiveCards(hand Cards)
	// SeeBoardState is called before players take action on Flop, Turn and River. It is also called once a
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
	// ActionUpdate lets the bot know of an action another bot player took, and the board state after the action
	// is complete.
	ActionUpdate(action Action, state BoardState)
}
```

Other relevant structures. You will be getting all of this information over GRPC.

```
type BoardState interface {
	// CommunityCards are the cards currently on the board. The first three cards will always be the flop, then the
	// turn, then the river.
	CommunityCards() Cards
	// Pot returns the size of the current Pot.
	Pot() int
	// Stage returns what the current Stage is, PreFlop, Flop, Turn or River.
	Stage() Stage
	// SmallBlindButton returns the index of the player in the Players slice that corresponds to the Small Blind Button.
	SmallBlindButton() int
	// Players return the array of all players in the game.
	Players() []PlayerState
}

type Action interface {
	// Type represents which action was taken play Player
	Type() ActionType
	// Player is which player took this action
	Player() PlayerState
	// Amount will be the amount of chips the ActionType refers too.
	Amount() int
}

type PlayerRoundResults interface {
   // ChipsWon returns how many chips this player won, 0 if none.
   ChipsWon() int
   // Cards will by nil if the player mucked, otherwise you will see their hand.
   Cards() Cards
   // HandStrength will return the engine's calcualtion of their hand strength.
   HandStrength() HandStrength
}

type PlayerState interface {
   // Stack returns how many chips this player has to bet.
   Stack() int
   // Status returns one of the possible PlayerStatus
   Status() PlayerStatus
   // SeatNumber return the index of this player on the board
   SeatNumber() int
   // Id returns a unique player ID.
   Id() string
   // PlayerRoundResults will be nil if this isnt a round end or if the player didn't make it there.
   PlayerRoundResults() PlayerRoundResults
}
```