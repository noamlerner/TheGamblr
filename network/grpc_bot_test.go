package network

import (
	"strconv"
	"testing"
	"time"

	"pokerengine/engine"
	"pokerengine/proto/thegamblr/proto"

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

	action, amount = bot.Act()
	assert.Equal(t, action, engine.RaiseAction)
	assert.Equal(t, amount, 20)

}

type flushUpdatesTest struct {
	name                 string
	isMyAction           bool
	includeBoardState    bool
	includeHand          bool
	includeActionUpdates bool
}

func TestFlushUpdates(t *testing.T) {
	tests := make([]*flushUpdatesTest, 0)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 2; l++ {
					tests = append(tests, &flushUpdatesTest{
						name:                 strconv.Itoa(i + 10*j + 100*k + 1000*l),
						isMyAction:           i == 0,
						includeBoardState:    j == 0,
						includeHand:          k == 0,
						includeActionUpdates: l == 0,
					})
				}
			}
		}
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bot := newGrpcBot()
			hand := engine.GenerateRandHand()[:2]
			boardState := engine.TestBoardState()
			if test.includeHand {
				bot.ReceiveCards(hand, 0, boardState)
			}

			if test.isMyAction {
				bot.isMyTurnToAct = true
			}

			if test.includeBoardState {
				bot.SeeBoardState(boardState)
			}

			if test.includeActionUpdates {
				bot.ActionUpdate(engine.NewAction(engine.CallAction, nil, 0))
				bot.ActionUpdate(engine.NewAction(engine.RaiseAction, nil, 10))
			}

			updates := bot.FlushUpdates()
			if test.includeHand {
				assert.Equal(t, protoConv.Cards(hand), updates.MyHand)
			} else {
				assert.Nil(t, updates.MyHand)

			}

			if test.isMyAction {
				assert.True(t, bot.isMyTurnToAct)
				assert.True(t, updates.IsMyAction)
			} else {
				assert.False(t, bot.isMyTurnToAct)
				assert.False(t, updates.IsMyAction)
			}

			if test.includeBoardState || test.includeHand {
				assert.Equal(t, protoConv.BoardState(boardState), updates.BoardState)
			}

			if test.includeActionUpdates {
				assert.Equal(t, protoConv.Actions([]engine.Action{
					engine.NewAction(engine.CallAction, nil, 0),
					engine.NewAction(engine.RaiseAction, nil, 10),
				}), updates.ActionUpdates)
			}
		})
	}
}
