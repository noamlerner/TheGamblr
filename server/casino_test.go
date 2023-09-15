package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/noamlerner/TheGamblr/proto/thegamblr/proto"

	"github.com/stretchr/testify/assert"
)

func TestGameFlow(t *testing.T) {
	casino := NewCasinoServer()
	ctx := context.Background()
	createGameRes, err := casino.CreateGame(ctx, &proto.CreateGameRequest{})
	assert.NoError(t, err)

	joinGameRes, err := casino.JoinGame(context.Background(), &proto.JoinGameRequest{
		PlayerId: "Bean",
		GameId:   createGameRes.GameId,
	})
	assert.NoError(t, err)
	beanToken := joinGameRes.Token

	joinGameRes, err = casino.JoinGame(context.Background(), &proto.JoinGameRequest{
		PlayerId: "Noam",
		GameId:   createGameRes.GameId,
	})
	assert.NoError(t, err)
	noamToken := joinGameRes.Token

	startGameResponse, err := casino.StartGame(context.Background(), &proto.StartGameRequest{
		Token: "randomtoken",
	})
	assert.NotNil(t, err)
	assert.Nil(t, startGameResponse)

	startGameResponse, err = casino.StartGame(context.Background(), &proto.StartGameRequest{
		Token: noamToken,
	})

	var updateResponse *proto.ReceiveUpdatesResponse
	for updateResponse.GetMyActionPacket() == nil {
		updateResponse, err = casino.ReceiveUpdates(context.Background(), &proto.ReceiveUpdatesRequest{
			Token: noamToken,
		})
		assert.NoError(t, err)
	}
	assert.Len(t, updateResponse.MyHand, 2)

	actionUpdates := updateResponse.Updates

	boardState := updateResponse.Updates[0].GetBoardState()
	assert.Equal(t, uint64(0), boardState.Pot)
	assert.Equal(t, proto.Stage_PRE_FLOP, boardState.Stage)
	assert.Equal(t, uint32(1), boardState.SmallBlindButton)
	assert.Equal(t, uint32(1), boardState.SmallBlindButton)
	assert.Equal(t, "Bean", boardState.Players[0].Id)
	assert.Equal(t, "Noam", boardState.Players[1].Id)

	assert.NotNil(t, updateResponse.MyActionPacket)
	assert.Equal(t, uint64(15), updateResponse.MyActionPacket.CurrentPot)
	assert.Equal(t, uint64(10), updateResponse.MyActionPacket.CallAmount)
	assert.Equal(t, uint64(5), updateResponse.MyActionPacket.LeftToCall)
	assert.Len(t, actionUpdates, 3)
	assert.Equal(t, uint64(5), actionUpdates[1].GetActionUpdate().Amount)
	assert.Equal(t, "Noam", actionUpdates[1].GetActionUpdate().Player.Id)
	assert.Equal(t, proto.ActionType_SMALL_BLIND, actionUpdates[1].GetActionUpdate().Type)

	assert.Equal(t, uint64(10), actionUpdates[2].GetActionUpdate().Amount)
	assert.Equal(t, "Bean", actionUpdates[2].GetActionUpdate().Player.Id)
	assert.Equal(t, proto.ActionType_BIG_BLIND, actionUpdates[2].GetActionUpdate().Type)

	updateResponse, err = casino.ReceiveUpdates(context.Background(), &proto.ReceiveUpdatesRequest{
		Token: beanToken,
	})
	assert.Nil(t, updateResponse.MyActionPacket)
	assert.True(t, reflect.DeepEqual(updateResponse.Updates, actionUpdates))

	updateResponse, err = casino.ReceiveUpdates(context.Background(), &proto.ReceiveUpdatesRequest{
		Token:          beanToken,
		SequenceNumber: updateResponse.NextSequenceNumber,
	})
	assert.NoError(t, err)
	assert.Len(t, updateResponse.Updates, 0)

	_, err = casino.Act(context.Background(), &proto.ActRequest{
		Token:      noamToken,
		ActionType: proto.ActionType_CALL,
	})
	assert.NoError(t, err)

	seqNum := updateResponse.NextSequenceNumber
	updateResponse = nil
	for updateResponse.GetMyActionPacket() == nil {
		updateResponse, err = casino.ReceiveUpdates(context.Background(), &proto.ReceiveUpdatesRequest{
			Token:          beanToken,
			SequenceNumber: seqNum,
		})
		assert.NoError(t, err)
	}

	assert.Len(t, updateResponse.Updates, 1)
	assert.Equal(t, "Noam", updateResponse.Updates[0].GetActionUpdate().Player.Id)
	assert.Equal(t, proto.ActionType_CALL, updateResponse.Updates[0].GetActionUpdate().Type)
	assert.NotNil(t, updateResponse.MyActionPacket)
	assert.Equal(t, uint64(20), updateResponse.MyActionPacket.CurrentPot)
	assert.Equal(t, uint64(10), updateResponse.MyActionPacket.CallAmount)
	assert.Equal(t, uint64(0), updateResponse.MyActionPacket.LeftToCall)

	_, err = casino.Act(ctx, &proto.ActRequest{
		Token:      beanToken,
		ActionType: proto.ActionType_CALL,
		Amount:     0,
	})
	assert.NoError(t, err)

	updateResponse.MyActionPacket = nil
	for updateResponse.GetMyActionPacket() == nil {
		updateResponse, err = casino.ReceiveUpdates(ctx, &proto.ReceiveUpdatesRequest{
			Token: noamToken,
		})
		assert.NoError(t, err)
	}
	_, err = casino.Act(ctx, &proto.ActRequest{
		Token:      noamToken,
		ActionType: proto.ActionType_RAISE,
		Amount:     10,
	})
	assert.NoError(t, err)

	seqNum = updateResponse.NextSequenceNumber
	updateResponse.MyActionPacket = nil
	for updateResponse.GetMyActionPacket() == nil {
		updateResponse, err = casino.ReceiveUpdates(ctx, &proto.ReceiveUpdatesRequest{
			Token:          beanToken,
			SequenceNumber: seqNum,
		})
		assert.NoError(t, err)
	}

	assert.NotNil(t, updateResponse)
	assert.Equal(t, "Noam", updateResponse.GetUpdates()[0].GetActionUpdate().Player.Id)
	assert.Equal(t, proto.ActionType_RAISE, updateResponse.GetUpdates()[0].GetActionUpdate().Type)
	assert.Equal(t, uint64(10), updateResponse.GetUpdates()[0].GetActionUpdate().Amount)
	assert.Len(t, updateResponse.GetUpdates(), 1)
	assert.Equal(t, uint64(10), updateResponse.GetMyActionPacket().GetCallAmount())
	assert.Equal(t, uint64(30), updateResponse.GetMyActionPacket().GetCurrentPot())
	assert.Equal(t, uint64(10), updateResponse.GetMyActionPacket().GetLeftToCall())

	_, err = casino.Act(ctx, &proto.ActRequest{
		Token:      beanToken,
		ActionType: proto.ActionType_FOLD,
	})
	assert.NoError(t, err)
	seqNum = updateResponse.NextSequenceNumber
	updateResponse = nil
	for len(updateResponse.GetUpdates()) < 2 {
		updateResponse, err = casino.ReceiveUpdates(ctx, &proto.ReceiveUpdatesRequest{
			Token:          beanToken,
			SequenceNumber: seqNum,
		})
		assert.NoError(t, err)
	}
	assert.NotNil(t, updateResponse.GetUpdates()[0].GetActionUpdate())
	assert.NotNil(t, updateResponse.GetUpdates()[1].GetBoardState())
}

func TestCasino_CreateGame(t *testing.T) {
	tests := []struct {
		name          string
		smallBLind    int
		numRounds     int
		startingStack int
	}{
		{
			"All defaults",
			0, 0, 0,
		},
		{
			"Semi-Custom",
			10, 200, 2000,
		},
		{
			"full custom",
			80, -1, 10000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			casino := NewCasinoServer()
			ctx := context.Background()
			createGame, err := casino.CreateGame(ctx, &proto.CreateGameRequest{
				SmallBlind:    uint64(test.smallBLind),
				NumRounds:     uint64(test.numRounds),
				StartingStack: uint64(test.startingStack),
			})
			assert.Nil(t, err)

			assert.NotEqual(t, createGame.GameId, "")
			game := casino.(*Casino).games[createGame.GameId]
			config := game.dealer.GameConfig()
			if test.smallBLind == 0 {
				assert.NotZero(t, config.SmallBlind)
			} else {
				assert.Equal(t, test.smallBLind, config.SmallBlind)
			}

			if test.startingStack == 0 {
				assert.NotZero(t, config.StartingStack)
			} else {
				assert.Equal(t, test.startingStack, config.StartingStack)
			}

			if test.numRounds == 0 {
				assert.NotZero(t, config.NumRounds)
			} else {
				assert.Equal(t, test.numRounds, config.NumRounds)
			}

		})
	}
}
