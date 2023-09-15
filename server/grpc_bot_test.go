package server

import (
	"testing"
	"time"

	"github.com/noamlerner/TheGamblr/engine"
	"github.com/noamlerner/TheGamblr/proto/thegamblr/proto"

	"github.com/stretchr/testify/assert"
)

func TestAct(t *testing.T) {
	bot := newGrpcBot()

	var action engine.ActionType
	var amount int

	time.AfterFunc(100*time.Millisecond, func() {
		bot.InputAction(&proto.ActRequest{
			ActionType: proto.ActionType(engine.RaiseAction),
			Amount:     20,
		})
		time.AfterFunc(100*time.Millisecond, func() {
			assert.Equal(t, action, engine.RaiseAction)
			assert.Equal(t, amount, 20)
		})
	})

	action, amount = bot.Act(0, 0, 0)
	assert.Equal(t, action, engine.RaiseAction)
	assert.Equal(t, amount, 20)

}

func TestFlushUpdates(t *testing.T) {

	bot := newGrpcBot()
	hand := engine.GenerateRandHand()[:2]
	boardState := engine.TestBoardState()
	bot.ReceiveCards(hand)
	bot.actionPacket = &proto.MyActionPacket{
		CurrentPot: 20,
		CallAmount: 20,
		LeftToCall: 10,
	}
	bot.SeeBoardState(boardState)
	bot.ActionUpdate(engine.NewAction(engine.CallAction, nil, 0))
	bot.ActionUpdate(engine.NewAction(engine.RaiseAction, nil, 10))
	updates := bot.FlushUpdates(0)
	assert.Equal(t, protoConv.Cards(hand), updates.MyHand)
	assert.NotNil(t, bot.actionPacket)
	assert.NotNil(t, updates.MyActionPacket)
	assert.Equal(t, protoConv.BoardState(boardState), updates.Updates[0].GetBoardState())
	assert.Equal(t, protoConv.Action(engine.NewAction(engine.CallAction, nil, 0)), updates.Updates[1].GetActionUpdate())
	assert.Equal(t, protoConv.Action(engine.NewAction(engine.RaiseAction, nil, 10)), updates.Updates[2].GetActionUpdate())
}
