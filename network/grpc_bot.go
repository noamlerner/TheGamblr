package network

import (
	"sync"

	"pokerengine/engine"
	"pokerengine/proto/thegamblr/proto"
)

type grpcBot struct {
	mutex         sync.Mutex
	cards         engine.Cards
	boardState    engine.BoardState
	actionUpdates []engine.Action
	isMyTurnToAct bool

	actionMutex   sync.Mutex
	actionChannel chan engine.Action
}

func newGrpcBot() *grpcBot {
	return &grpcBot{
		mutex:         sync.Mutex{},
		actionMutex:   sync.Mutex{},
		actionChannel: make(chan engine.Action),
	}
}
func (g *grpcBot) ReceiveCards(hand engine.Cards, blind int, boardState engine.BoardState) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.cards = hand
	g.boardState = boardState
}

func (g *grpcBot) SeeBoardState(boardState engine.BoardState) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.boardState = boardState
}

func (g *grpcBot) Act() (engine.ActionType, int) {
	g.turnOnMyTurnToAct()
	var action engine.Action
	select {
	case action = <-g.actionChannel:
		break
	}
	g.turnOffMyTurnToAct()
	return action.Type(), action.Amount()
}

func (g *grpcBot) turnOffMyTurnToAct() {
	g.actionMutex.Lock()
	defer g.actionMutex.Unlock()
	g.isMyTurnToAct = false
}

func (g *grpcBot) turnOnMyTurnToAct() {
	g.actionMutex.Lock()
	defer g.actionMutex.Unlock()
	g.isMyTurnToAct = true
}

func (g *grpcBot) InputAction(request *proto.ActRequest) {
	g.actionMutex.Lock()
	defer g.actionMutex.Unlock()
	if !g.isMyTurnToAct {
		return
	}
	g.actionChannel <- engine.NewAction(engine.ActionType(request.ActionType), nil, int(request.Amount))
}

func (g *grpcBot) ActionUpdate(action engine.Action) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.actionUpdates = append(g.actionUpdates, action)
}

func (g *grpcBot) FlushUpdates() *proto.ReceiveUpdatesResponse {
	g.mutex.Lock()
	defer g.mutex.Unlock()
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
