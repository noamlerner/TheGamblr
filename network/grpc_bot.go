package network

import (
	"pokerengine/engine"
	"pokerengine/proto/thegamblr/proto"
)

type grpcBot struct {
	token         string
	cards         engine.Cards
	boardState    engine.BoardState
	actionUpdates []engine.Action
	isMyTurnToAct bool

	actionChannel chan engine.Action
}

func (g *grpcBot) ReceiveCards(hand engine.Cards, blind int, boardState engine.BoardState) {
	g.cards = hand
	g.boardState = boardState
}

func (g *grpcBot) SeeBoardState(boardState engine.BoardState) {
	g.boardState = boardState
}

func (g *grpcBot) Act() (engine.ActionType, int) {
	g.isMyTurnToAct = true
	action := <-g.actionChannel
	g.isMyTurnToAct = false
	return action.Type(), action.Amount()
}

func (g *grpcBot) ActionUpdate(action engine.Action) {
	g.actionUpdates = append(g.actionUpdates, action)
}

func (g *grpcBot) InputAction(request *proto.ActRequest) {
	if !g.isMyTurnToAct {
		return
	}
	g.actionChannel <- engine.NewAction(engine.ActionType(request.ActionType), nil, int(request.Amount))
}

func (g *grpcBot) FlushUpdates() *proto.ReceiveUpdatesResponse {
	update := &proto.ReceiveUpdatesResponse{
		ActionUpdates: protoConv.Actions(g.actionUpdates),
		IsMyAction:    g.isMyTurnToAct,
		BoardState:    protoConv.BoardState(g.boardState),
		MyHand:        protoConv.Cards(g.cards),
	}

	g.cards = nil
	g.boardState = nil
	g.actionUpdates = nil
	return update
}
