# Engine

This is TheGamblr's engine package. This knows how to run a game and it can also be used for bot development in Go.

First lets give you an overview of what is in this package

## Overview

Most of this package is commented, so I'll just highlight some of the more useful structs and funcs you should know
about.
For Usage information, look at the comments in the code.

- [Deck](./deck.go) creates a deck for you that can be shuffled and used for dealing. There's also
  a [HardOrderedDeck](./hardordered_deck.go) which can be used to simulate games where certain hands and cards are
  dealt.
- [Dealer](./dealer.go) Managed the [board](./board.go) and players. It runs the game.
- The [BotPlayer](./player.go) interface is the code that needs to be defined for your player to play poker. We have a
  few simple ones implemented if you want to try them out including
    - [OneActionBot](./one_action_bot.go) which will always take one action
    - [SequenceBot](./sequence_of_actions_bot.go) which will take a sequence of actions, then start calling
    - [RandomActionBot](./random_action_bot.go) which will mostly call, but sometimes raise and rarely fold.
- [GameSim](./game_sim.go) which will let you run many games on a certain set of bots to let you know which ones perform
  the best
- [OddsOfWinning](./odds_of_winning.go) provides a function that tells you the odds of winning (or tieing) a game. You
  can specify your hand, the hands of other players, which community cards are dealt, and how many players are in the
  game.

## QuickStart Guide

Let's write some code using this package that sets two bots against each other. All this code can be found in
the quickstart package.

If you haven't yet, you will need to run

```
 go get github.com/noamlerner/TheGamblr
```

We need to start by defining our bot according to the [BotPlayer](./player.go) interface. Let's create one
that always calls to keep it simple.

```

import "github.com/noamlerner/TheGamblr/engine"

type OurBot struct {
}

func NewOurBot() *OurBot {
	return &OurBot{}
}

func (c *OurBot) ReceiveCards(hand engine.Cards) {
}

func (c *OurBot) SeeBoardState(boardState engine.BoardState) {
}

func (c *OurBot) Act(int, int, int) (engine.ActionType, int) {
	return engine.CallAction, 10
}

func (c *OurBot) ActionUpdate(action engine.Action, state engine.BoardState) {
}
```

That was easy! Our bot is just always going to call. Let's create a game between two of these bots and run it once.
We'll start by creating a `main` package with a `main` file and a `main` func.

Inside the func we will set the two bots to play against each other by using the `GameSim` struct. The `GameSim`
constructor takes in `BotPlayerProviders`, which returns the player ID and the BotPlayer, so lets start by creating
those.

```
	playerProviders := []engine.BotPlayerProvider{
		func() (string, engine.BotPlayer) {
			return "TheGamblr", engine.NewRandomActionBot()
		},
		func() (string, engine.BotPlayer) {
			return "Bean", quickstart.NewOurBot()
		},
	}
```

There we have two functions. The first returns "TheGamblr" which uses the `RandomActionBot` provided by the engine.
The second returns "Bean" (named after my dog) which uses the bot we just created!

We can use this to construct our `GameSim` and set it to only run 1 game.

```
		gameSim := engine.NewGameSim(playerProviders).WithNumSims(1)
```

Now we will create a custom game config to have some interesting output.

```
	gameConfig := &engine.GameConfig{
		// SmallBlind will be 5
		SmallBlind:    5,
		// They only have 10 rounds until the game is over
		NumRounds:     10,
		// each player starts with 100 chips
		StartingStack: 100,
		// The highest log level, everything will be outputted.
		LogLevel:      engine.LogLevelCards,
	}
```

Now we pass the `gameConfig` into the `gameSim` and run to see our two bots play against each other!

```
	gameSim.WithGameConfig(gameConfig).Run()
```

That's it! We ran our first simulation.

Check the output of the first 3 rounds:

```
Round 1
Bean Pays Small Blind 5
Bean received 8 of Hearts and T of Clubs
TheGamblr Pays Big Blind 10
TheGamblr received 7 of Clubs and A of Hearts
Bean Calls 5
TheGamblr Raises To 20
Bean Calls 10
Flop:
        CommunityCards: 2 of Hearts, 3 of Clubs and 3 of Spades,
        Pot 40
Bean Checks
TheGamblr Checks
Turn:
        CommunityCards: 2 of Hearts, 3 of Clubs, 3 of Spades and A of Spades,
        Pot 40
Bean Checks
TheGamblr Checks
River:
        CommunityCards: 2 of Hearts, 3 of Clubs, 3 of Spades, A of Spades and Q of Hearts,
        Pot 40
Bean Checks
TheGamblr Checks
TheGamblr won 40 chips with 7 of Clubs and A of Hearts making a TwoPair
Stack Sizes:
        TheGamblr - 120
        Bean - 80
Round 2
TheGamblr Pays Small Blind 5
TheGamblr received T of Clubs and T of Diamonds
Bean Pays Big Blind 10
Bean received 9 of Spades and 8 of Diamonds
TheGamblr Calls 5
Bean Checks
Flop:
        CommunityCards: J of Diamonds, T of Hearts and 6 of Hearts,
        Pot 20
TheGamblr Checks
Bean Checks
Turn:
        CommunityCards: J of Diamonds, T of Hearts, 6 of Hearts and 6 of Clubs,
        Pot 20
TheGamblr Folds
Stack Sizes:
        TheGamblr - 110
        Bean - 90
Round 3
Bean Pays Small Blind 5
Bean received 6 of Hearts and 6 of Diamonds
TheGamblr Pays Big Blind 10
TheGamblr received 5 of Clubs and 7 of Spades
Bean Calls 5
TheGamblr Raises To 20
Bean Calls 10
Flop:
        CommunityCards: J of Spades, 8 of Diamonds and A of Hearts,
        Pot 40
Bean Checks
TheGamblr Checks
Turn:
        CommunityCards: J of Spades, 8 of Diamonds, A of Hearts and 2 of Diamonds,
        Pot 40
Bean Checks
TheGamblr Checks
River:
        CommunityCards: J of Spades, 8 of Diamonds, A of Hearts, 2 of Diamonds and 9 of Spades,
        Pot 40
Bean Checks
TheGamblr Checks
TheGamblr won 0 chips with 5 of Clubs and 7 of Spades making a HighCard
Bean won 40 chips with 6 of Hearts and 6 of Diamonds making a Pair
Stack Sizes:
        TheGamblr - 90
        Bean - 110
```
