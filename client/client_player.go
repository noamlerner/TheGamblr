package client

import (
	"context"
	"fmt"

	"pokerengine/engine"
	"pokerengine/proto/thegamblr/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientBot struct {
	actor  engine.BotPlayer
	casino proto.CasinoClient

	name  string
	token string
}

func NewClientBot(name string, actor engine.BotPlayer, connectionString string) *ClientBot {
	conn, err := grpc.Dial(connectionString, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &ClientBot{actor: actor, casino: proto.NewCasinoClient(conn), name: name}
}

func (c *ClientBot) JoinGame(ctx context.Context, gameID string) {
	res, err := c.casino.JoinGame(ctx, &proto.JoinGameRequest{GameId: gameID, PlayerId: c.name})
	if err != nil {
		panic(err)
	}
	c.token = res.Token
}

func (c *ClientBot) StartGame(ctx context.Context) {
	_, err := c.casino.StartGame(ctx, &proto.StartGameRequest{Token: c.token})
	if err != nil {
		panic(err)
	}
}
func (c *ClientBot) CreateGame(ctx context.Context) string {
	res, err := c.casino.CreateGame(ctx, &proto.CreateGameRequest{})
	if err != nil {
		panic(err)
	}
	return res.GameId
}

func (c *ClientBot) Run(ctx context.Context) {
	seq_num := uint64(0)
	gameOver := false
	for !gameOver {
		seq_num, gameOver = c.updateLoop(ctx, seq_num)
	}
}

func (c *ClientBot) updateLoop(ctx context.Context, seq_num uint64) (uint64, bool) {
	updateRes, _ := c.casino.ReceiveUpdates(ctx, &proto.ReceiveUpdatesRequest{
		Token:          c.token,
		SequenceNumber: seq_num,
	})

	for _, update := range updateRes.GetUpdates() {
		if c.informOfUpdate(update) {
			return seq_num, true
		}
	}

	if updateRes.GetMyActionPacket() != nil {
		c.act(ctx, updateRes)
	}
	return updateRes.GetNextSequenceNumber(), false
}

func (c *ClientBot) informOfUpdate(update *proto.Update) bool {
	action := update.GetActionUpdate()
	if action != nil {
		player := action.Player
		c.actor.ActionUpdate(engine.NewAction(engine.ActionType(action.Type), convertProtoPlayer(player), int(action.Amount)))
	} else {
		board := update.GetBoardState()
		players := make([]engine.PlayerState, len(board.GetPlayers()))
		for i, player := range board.GetPlayers() {
			players[i] = convertProtoPlayer(player)
		}
		boardState := engine.NewBoardState(convertProtoCards(board.CommunityCards), int(board.Pot), engine.Stage(board.Stage), int(board.SmallBlindButton), players)
		c.actor.SeeBoardState(boardState)
		if boardState.Stage() == engine.GameOver {
			return true
		}
	}
	return false
}

func convertProtoPlayer(player *proto.PlayerState) engine.PlayerState {
	if player == nil {
		return nil
	}
	var roundResults engine.PlayerRoundResults
	if player.GetRoundResults() != nil {
		pbResults := player.RoundResults
		cards := convertProtoCards(pbResults.Cards)
		roundResults = engine.NewPlayerRoundResults(int(pbResults.ChipsWon), cards, engine.HandStrength(pbResults.HandStrength))
	}

	p := engine.NewPlayerState(int(player.GetStack()), engine.PlayerStatus(player.GetStatus()), int(player.GetSeatNumber()), player.GetId(), roundResults)
	return p
}

func convertProtoCards(pbCards []*proto.Card) engine.Cards {
	cards := make(engine.Cards, len(pbCards))
	for i, card := range pbCards {
		cards[i] = engine.NewCard(engine.Rank(card.Rank), engine.Suit(card.Suit))
	}
	return cards
}

func (c *ClientBot) act(ctx context.Context, updateRes *proto.ReceiveUpdatesResponse) {
	packet := updateRes.MyActionPacket
	action, amount := c.actor.Act(int(packet.CurrentPot), int(packet.CallAmount), int(packet.LeftToCall))

	err := fmt.Errorf("")
	for err != nil {
		_, err = c.casino.Act(ctx, &proto.ActRequest{
			Token:      c.token,
			ActionType: proto.ActionType(action),
			Amount:     int64(amount),
		})
	}
}
