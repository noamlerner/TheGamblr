package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var names = []string{"Noam", "Bean", "TheGamblr", "Bob", "Lulu", "Shlomi", "Lamp", "Leica"}

func TestFullGame_3Players(t *testing.T) {
	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelCards
	config.NumRounds = 4
	dealer := NewDealer(config)
	dealer.SeatPlayer("Noam", NewRandomActionBot())
	dealer.SeatPlayer("Bean", NewRandomActionBot())
	dealer.SeatPlayer("TheGamblr", NewRandomActionBot())
	dealer.RunGame()
}

func TestFullGame_2Players(t *testing.T) {
	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelCards
	dealer := NewDealer(config)
	dealer.SeatPlayer("Noam", NewRandomActionBot())
	dealer.SeatPlayer("Bean", NewRandomActionBot())
	dealer.RunGame()
}

func TestSequence_A(t *testing.T) {
	numPlayers := 3
	sequence := []ActionType{RaiseAction, RaiseAction, RaiseAction, CallAction}
	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelCards
	config.NumRounds = 10
	d := NewDealer(config)
	for i := 0; i < numPlayers; i++ {
		d.SeatPlayer(names[i], &SequenceOfActionsBot{})
	}

	for i, action := range sequence {
		bot := d.board.players[i%numPlayers].actor.(*SequenceOfActionsBot)
		bot.sequence = append(bot.sequence, action)
	}
	d.board.smallBlindButton = 3
	d.newRound()
	// PreFlop betting
	d.betting()
	assert.Equal(t, 2, d.board.players[0].actor.(*SequenceOfActionsBot).numCalled)
	assert.Equal(t, 1, d.board.players[1].actor.(*SequenceOfActionsBot).numCalled)
	assert.Equal(t, 2, d.board.players[2].actor.(*SequenceOfActionsBot).numCalled)
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
	d := NewDealer(config)
	for i := 0; i < numPlayers; i++ {
		d.SeatPlayer(names[i], &OneActionBot{action: CallAction})
	}
	d.board.smallBlindButton = 3

	d.deck = NewHardOrderedDeck(hands, numPlayers, cards)
	d.playRound()
}

func TestSequence_AllInLowStackWins(t *testing.T) {
	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelCards
	config.NumRounds = 10
	d := NewDealer(config)
	d.SeatPlayer(names[0], &OneActionBot{action: CallAction})
	d.SeatPlayer(names[1], &OneActionBot{action: RaiseAction})
	d.SeatPlayer(names[2], &OneActionBot{action: CallAction})

	d.board.players[0].stack = 12
	d.board.players[1].stack = 15
	d.board.players[2].stack = 12

	hands := []Cards{
		// hand 0
		{NewCard(Eight, Spades), NewCard(Six, Spades)},
		// hand 1
		{NewCard(Queen, Diamonds), NewCard(Ace, Clubs)},
		// hand 2
		{NewCard(Five, Clubs), NewCard(Six, Diamonds)},
	}
	cards := Cards{
		// community cards
		NewCard(Ten, Spades), NewCard(Four, Spades), NewCard(Seven, Spades), NewCard(Ace, Hearts), NewCard(Ace, Spades),
	}
	d.deck = NewHardOrderedDeck(hands, 3, cards)
	d.board.smallBlindButton = 3
	d.playRound()

	// Everyone calls, they should only be able to win 12. player 1 should be getting back 3 chips.
	assert.Equal(t, 3, d.board.players[1].stack)
}

func TestSequence_EveryoneFolds(t *testing.T) {
	sequence := []ActionType{CallAction, FoldAction}
	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelCards
	config.NumRounds = 10
	d := NewDealer(config)
	d.SeatPlayer(names[0], &SequenceOfActionsBot{sequence: sequence})
	d.SeatPlayer(names[1], &SequenceOfActionsBot{sequence: sequence})
	d.SeatPlayer(names[2], &SequenceOfActionsBot{})

	d.playRound()
	assert.Equal(t, 990, d.board.players[0].stack)
	assert.Equal(t, 990, d.board.players[1].stack)
	assert.Equal(t, 1020, d.board.players[2].stack)
}
