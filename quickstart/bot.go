package quickstart

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
