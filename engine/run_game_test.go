package engine

import (
	"testing"
)

var names = []string{"Noam", "Bean", "TheGamblr"}

func TestFullGame(t *testing.T) {
	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelCards
	config.NumRounds = 4
	dealer := Dealer(config)
	dealer.SeatPlayer("Noam", NewRandomActionBot())
	dealer.SeatPlayer("Bean", NewRandomActionBot())
	dealer.SeatPlayer("TheGamblr", NewRandomActionBot())
	dealer.RunGame()
}

func TestSequence(t *testing.T) {
	numPlayers := 3
	sequence := []Action{RaiseAction, RaiseAction, RaiseAction, CallAction}
	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelCards
	config.NumRounds = 10
	d := Dealer(config)
	for i := 0; i < numPlayers; i++ {
		d.SeatPlayer(names[i], &SequenceOfActionsBot{})
	}

	for i, action := range sequence {
		bot := d.board.players[i%numPlayers].actor.(*SequenceOfActionsBot)
		bot.actions = append(bot.actions, action)
	}
	d.board.smallBlindButton = 3
	d.playRound()
}

func TestHardcodedCards(t *testing.T) {
	numPlayers := 3
	hands := []Cards{
		// hand 0
		{NewCard(Eight, Spades), NewCard(Six, Clubs)},
		// hand 1
		{NewCard(Queen, Diamonds), NewCard(Ace, Clubs)},
		// hand 2
		{NewCard(Five, Clubs), NewCard(Six, Spades)},
	}
	cards := Cards{
		// community cards
		NewCard(Ten, Spades), NewCard(Four, Spades), NewCard(Seven, Spades), NewCard(Ace, Hearts), NewCard(Ace, Spades),
	}

	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelCards
	config.NumRounds = 10
	d := Dealer(config)
	for i := 0; i < numPlayers; i++ {
		d.SeatPlayer(names[i], &OneActionBot{action: CallAction})
	}
	d.board.smallBlindButton = 3

	d.deck = NewHardOrderedDeck(hands, numPlayers, cards)
	d.playRound()
}
