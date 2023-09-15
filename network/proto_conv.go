package network

import (
	"pokerengine/engine"
	"pokerengine/proto/thegamblr/proto"
)

type _protoConv struct{}

var protoConv = &_protoConv{}

func (c *_protoConv) Card(card *engine.Card) *proto.Card {
	return &proto.Card{
		Rank: proto.Rank(card.Rank()),
		Suit: proto.Suit(card.Suit()),
	}
}

func (c *_protoConv) Cards(cards engine.Cards) []*proto.Card {
	if cards == nil {
		return nil
	}
	pCards := make([]*proto.Card, len(cards))
	for i, card := range cards {
		pCards[i] = c.Card(card)
	}
	return pCards
}

func (c *_protoConv) BoardState(b engine.BoardState) *proto.BoardState {
	if b == nil {
		return nil
	}
	return &proto.BoardState{
		CommunityCards:   c.Cards(b.CommunityCards()),
		Pot:              uint64(b.Pot()),
		Stage:            proto.Stage(b.Stage()),
		SmallBlindButton: uint32(b.SmallBlindButton()),
		Players:          c.Players(b.Players()),
	}
}

func (c *_protoConv) Player(p engine.PlayerState) *proto.PlayerState {
	return &proto.PlayerState{
		Stack:        uint64(p.Stack()),
		Status:       proto.PlayerStatus(p.Status()),
		SeatNumber:   uint32(p.SeatNumber()),
		Id:           p.Id(),
		RoundResults: c.RoundResults(p.RoundEndStats()),
	}
}

func (c *_protoConv) Players(players []engine.PlayerState) []*proto.PlayerState {
	pPlayers := make([]*proto.PlayerState, len(players))
	for i, player := range players {
		pPlayers[i] = c.Player(player)
	}
	return pPlayers
}

func (c *_protoConv) RoundResults(p engine.PlayerRoundResults) *proto.PlayerRoundResults {
	return &proto.PlayerRoundResults{
		ChipsWon:     uint64(p.ChipsWon()),
		Cards:        c.Cards(p.Cards()),
		HandStrength: proto.HandStrength(p.HandStrength()),
	}
}

func (c *_protoConv) Action(a engine.Action) *proto.Action {
	return &proto.Action{
		Type:   proto.ActionType(a.Type()),
		Player: c.Player(a.Player()),
		Amount: uint64(a.Amount()),
	}
}

func (c *_protoConv) Actions(actions []engine.Action) []*proto.Action {
	if actions == nil {
		return nil
	}
	pActions := make([]*proto.Action, len(actions))
	for i, action := range actions {
		pActions[i] = c.Action(action)
	}
	return pActions
}
