package client

import (
	"github.com/noamlerner/TheGamblr/engine"
	"github.com/noamlerner/TheGamblr/proto/thegamblr/proto"
)

type _protoConv struct{}

var protoConv = &_protoConv{}

func (p *_protoConv) convertProtoPlayer(player *proto.PlayerState) engine.PlayerState {
	if player == nil {
		return nil
	}
	var roundResults engine.PlayerRoundResults
	if player.GetRoundResults() != nil {
		pbResults := player.RoundResults
		cards := p.convertProtoCards(pbResults.Cards)
		roundResults = engine.NewPlayerRoundResults(int(pbResults.ChipsWon), cards, engine.HandStrength(pbResults.HandStrength))
	}

	return engine.NewPlayerState(int(player.GetStack()), engine.PlayerStatus(player.GetStatus()), int(player.GetSeatNumber()), player.GetId(), roundResults)
}

func (p *_protoConv) convertProtoCards(pbCards []*proto.Card) engine.Cards {
	cards := make(engine.Cards, len(pbCards))
	for i, card := range pbCards {
		cards[i] = engine.NewCard(engine.Rank(card.Rank), engine.Suit(card.Suit))
	}
	return cards
}

func (p *_protoConv) convertBoard(board *proto.BoardState) engine.BoardState {
	players := make([]engine.PlayerState, len(board.GetPlayers()))
	for i, player := range board.GetPlayers() {
		players[i] = p.convertProtoPlayer(player)
	}
	boardState := engine.NewBoardState(p.convertProtoCards(board.CommunityCards), int(board.Pot), engine.Stage(board.Stage), int(board.SmallBlindButton), players)
	return boardState
}
