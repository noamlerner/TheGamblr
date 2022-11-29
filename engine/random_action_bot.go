package engine

import (
	"math/rand"
	"time"
)

func NewRandomActionBot() BotPlayer {
	return &randomActionBot{
		rand: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

type randomActionBot struct {
	rand *rand.Rand
}

func (c *randomActionBot) RoundResults(results RoundResults) {
}

func (c *randomActionBot) ReceiveCards(hand Cards, blind int, boardState ActiveBoard) {
}

func (c *randomActionBot) SeeActiveBoardState(boardState ActiveBoard) {
}

func (c *randomActionBot) Act() (Action, int) {
	f := c.rand.Float64()
	switch {
	case f < 0.5:
		return CallAction, 10
	case f >= 0.5 && f < 0.95:
		return RaiseAction, 10
	case f >= 0.95:
		return FoldAction, 10
	}
	return CheckFoldAction, 0
}

func (c *randomActionBot) ReceiveUpdate(action VisibleAction) {
}
