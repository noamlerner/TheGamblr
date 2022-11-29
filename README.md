# The Gamblr

_Legend has it that this entire codebase was written in one playthrough
of [The Gambler by Kenny Rogers](https://www.youtube.com/watch?v=7hx4gdlfamo)_

...Those who know the author know he probably just put it on repeat for the whole time.


## Intro

_You gotta know when to hold 'em, know when to fold 'em_

This is a poker engine written to allow people to pit algorithms against each other. While this engine is written in Go,
it allows its users to use any language they want by creating an HTTP interface through which the game will be
played.
        HTTP Interface still in progress


## Available Information

_I've made a life out of reading people's faces_

In order to make decisions, you're going need to know your input.

While the names are still subject to change, here is the information that will be available to you.

### ActiveBoardState

_You never count your money when you're sittin' at the table_

When you get your cards and then again after everytime CommunityCards are added, you will receive the ActiveBoard state.

```
type ActiveBoard interface {
        // CommunityCards are the cards currently on the board. The first three cards will always be the flop, then the
        // turn, then the river.
        CommunityCards() []*Card
        // Pot returns the size of the current Pot.
        Pot() int
        // Stage returns what the current Stage is, PreFlop, Flop, Turn or River.
        Stage() Stage
        // SmallBlindButton returns the index of the player in the Players slice that corresponds to the Small Blind Button.
        SmallBlindButton() int
        // Players return the array of all players in the game.
        Players() []ActivePlayerState
}

type ActivePlayerState interface {
        // Stack returns how many chips this player has to bet.
        Stack() int
        // Status returns one of the possible BoardPlayerStatus
        Status() BoardPlayerStatus
        // SeatNumber return the index of this player on the board
        SeatNumber() int
        // Id returns a unique player ID.
        Id() string
}
const (
        // BoardPlayerStatusOut means the player has lost all their money
        BoardPlayerStatusOut BoardPlayerStatus = iota
        // BoardPlayerStatusFolded means the player has folded the round
        BoardPlayerStatusFolded
        // BoardPlayerStatusPlaying means the player is in the round
        BoardPlayerStatusPlaying
        // BoardPlayerStatusAllIn is for players that are in the game, but are currently AllIn
        BoardPlayerStatusAllIn
)

```

### Mid-Game Actions
_Boredom overtook us and he began to speak_

During the betting rounds, you will receive `ActionUpdates` Every time an action is taken.
```
type VisibleAction interface {
        // ActionTaken represents which action was taken play Player
        ActionTaken() Action
        // Player is which player took this action
        Player() ActivePlayerState
        // Amount will be the amount of chips the Action refers too.
        Amount() int
}

FoldAction Action = iota
CallAction
RaiseAction
SmallBlind
BigBlind
```

### Round Results

_In his final words I found an ace the I could keep_

At the end of each round you will get the RoundResults which are similar to the active board. It contains information
on any hands that weren't mucked as well as who won chips and how many.

```
type RoundResults interface {
        CommunityCards() []*Card
        Pot() int
        SmallBlindButton() int
        PlayerResults() []PlayerResults
}

type PlayerResults interface {
        // Stack returns how many chips this player has to bet.
        Stack() int
        // Status returns one of the possible BoardPlayerStatus
        Status() BoardPlayerStatus
        // SeatNumber return the index of this player on the board
        SeatNumber() int
        // Id returns a unique player ID.
        Id() string
        // ChipsWon returns how many chips this player won, 0 if none.
        ChipsWon() int
        // Cards will by nil if the player mucked, otherwise you will see their hand.
        Cards() Cards
        // HandStrength will return the engine's calcualtion of their hand strength.
        HandStrength() HandStrength
}

const (
        HandStrengthUnset HandStrength = iota
        HighCard
        Pair
        TwoPair
        Trips
        Straight
        Flush
        FullHouse
        Quads
        StraightFlush
)
```

### Your Algorithm

_If you're gonna play the game, boy you gotta learn to play it right_

You should expect your bot to implement this interface

```
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
```