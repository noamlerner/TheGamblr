package client

import (
	"context"
	"fmt"

	"github.com/noamlerner/TheGamblr/engine"
	"github.com/noamlerner/TheGamblr/proto/thegamblr/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCBot struct {
	actor  engine.BotPlayer
	casino proto.CasinoClient

	name  string
	token string
}

func NewClientBot(name string, actor engine.BotPlayer, connectionString string) *GRPCBot {
	conn, err := grpc.Dial(connectionString, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &GRPCBot{actor: actor, casino: proto.NewCasinoClient(conn), name: name}
}

func (c *GRPCBot) JoinGame(ctx context.Context, gameID string) {
	res, err := c.casino.JoinGame(ctx, &proto.JoinGameRequest{GameId: gameID, PlayerId: c.name})
	if err != nil {
		panic(err)
	}
	c.token = res.Token
}

func (c *GRPCBot) StartGame(ctx context.Context) {
	_, err := c.casino.StartGame(ctx, &proto.StartGameRequest{Token: c.token})
	if err != nil {
		panic(err)
	}
}
func (c *GRPCBot) CreateGame(ctx context.Context) string {
	res, err := c.casino.CreateGame(ctx, &proto.CreateGameRequest{})
	if err != nil {
		panic(err)
	}
	return res.GameId
}

func (c *GRPCBot) Run(ctx context.Context) {
	seqNum := uint64(0)
	gameOver := false
	for !gameOver {
		// keep running the update loop until the game is over. We update the seqNum at each iteration with the one
		// returned to us by the Casino.
		seqNum, gameOver = c.updateLoop(ctx, seqNum)
	}
}

func (c *GRPCBot) updateLoop(ctx context.Context, seqNum uint64) (uint64, bool) {
	// any new updates will be here.
	updateRes, _ := c.casino.ReceiveUpdates(ctx, &proto.ReceiveUpdatesRequest{
		Token:          c.token,
		SequenceNumber: seqNum,
	})

	// pass the updates along to our bot
	for _, update := range updateRes.GetUpdates() {
		// informUpdate will return true if the game is over
		if c.informOfUpdate(updateRes, update) {
			return seqNum, true
		}
	}

	// if it's my turn, it's time to act.
	if updateRes.GetMyActionPacket() != nil {
		c.act(ctx, updateRes)
	}
	return updateRes.GetNextSequenceNumber(), false
}

func (c *GRPCBot) informOfUpdate(res *proto.ReceiveUpdatesResponse, update *proto.Update) bool {
	// Update can either be an ActionUpdate or a BoardState.
	action := update.GetActionUpdate()
	if action != nil {
		player := action.Player
		c.actor.ActionUpdate(engine.NewAction(engine.ActionType(action.Type), protoConv.convertProtoPlayer(player), int(action.Amount)))
	} else {
		boardState := protoConv.convertBoard(update.GetBoardState())
		if boardState.Stage() == engine.PreFlop {
			c.actor.ReceiveCards(protoConv.convertProtoCards(res.GetMyHand()))
		}
		c.actor.SeeBoardState(boardState)
		if boardState.Stage() == engine.GameOver {
			return true
		}
	}
	return false
}

func (c *GRPCBot) act(ctx context.Context, updateRes *proto.ReceiveUpdatesResponse) {
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
