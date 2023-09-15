package server

import (
	"sync"

	"github.com/noamlerner/TheGamblr/engine"
	"github.com/noamlerner/TheGamblr/proto/thegamblr/proto"
)

type grpcBot struct {
	mutex        sync.RWMutex
	cards        engine.Cards
	updates      []*proto.Update
	actionPacket *proto.MyActionPacket

	actionMutex   sync.Mutex
	actionChannel chan engine.Action
}

func newGrpcBot() *grpcBot {
	return &grpcBot{
		mutex:         sync.RWMutex{},
		actionMutex:   sync.Mutex{},
		actionChannel: make(chan engine.Action),
	}
}
func (g *grpcBot) ReceiveCards(hand engine.Cards) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.cards = hand
}

func (g *grpcBot) SeeBoardState(boardState engine.BoardState) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.updates = append(g.updates, &proto.Update{
		Update: &proto.Update_BoardState{
			BoardState: protoConv.BoardState(boardState),
		},
		SequenceNumber: uint64(len(g.updates)),
	})
}

func (g *grpcBot) Act(pot, amountToCall, leftToCall int) (engine.ActionType, int) {
	g.setActionPacket(pot, amountToCall, leftToCall)
	var action engine.Action
	select {
	case action = <-g.actionChannel:
		break
	}
	g.unsetActionPacket()
	return action.Type(), action.Amount()
}

func (g *grpcBot) unsetActionPacket() {
	g.actionMutex.Lock()
	defer g.actionMutex.Unlock()
	g.actionPacket = nil
}

func (g *grpcBot) setActionPacket(pot, amountToCall, leftToCall int) {
	g.actionMutex.Lock()
	defer g.actionMutex.Unlock()
	g.actionPacket = &proto.MyActionPacket{
		CurrentPot: uint64(pot),
		CallAmount: uint64(amountToCall),
		LeftToCall: uint64(leftToCall),
	}
}

func (g *grpcBot) InputAction(request *proto.ActRequest) {
	g.actionMutex.Lock()
	defer g.actionMutex.Unlock()
	if g.actionPacket == nil {
		return
	}
	g.actionChannel <- engine.NewAction(engine.ActionType(request.ActionType), nil, int(request.Amount))
}

func (g *grpcBot) ActionUpdate(action engine.Action) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.updates = append(g.updates, &proto.Update{
		Update:         &proto.Update_ActionUpdate{ActionUpdate: protoConv.Action(action)},
		SequenceNumber: uint64(len(g.updates)),
	})
}

func (g *grpcBot) FlushUpdates(prevSequenceNumber int) *proto.ReceiveUpdatesResponse {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	var updates []*proto.Update
	if prevSequenceNumber < len(g.updates) {
		updates = g.updates[prevSequenceNumber:]
	}
	return &proto.ReceiveUpdatesResponse{
		Updates:            updates,
		MyActionPacket:     g.actionPacket,
		MyHand:             protoConv.Cards(g.cards),
		NextSequenceNumber: uint64(len(g.updates)),
	}
}
