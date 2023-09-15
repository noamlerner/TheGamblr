package network

import (
	context "context"
	"fmt"
	"strings"

	"pokerengine/engine"
	"pokerengine/proto/thegamblr/proto"

	"github.com/google/uuid"
)

const (
	tokenGameDelimiter = "~"
	tokenFormat        = "%s~%s" // gameID~token
)

type Casino struct {
	proto.UnimplementedCasinoServer
	games map[string]*game
}

func NewCasinoServer() proto.CasinoServer {
	return &Casino{
		games: map[string]*game{},
	}
}

func (c *Casino) CreateGame(ctx context.Context, request *proto.CreateGameRequest) (*proto.CreateGameResponse, error) {
	gameID := uuid.Must(uuid.NewRandom()).String()

	config := &engine.GameConfig{
		SmallBlind:    int(request.SmallBlind),
		NumRounds:     int(request.NumRounds),
		StartingStack: int(request.StartingStack),
	}

	if config.SmallBlind == 0 {
		config.SmallBlind = 5
	}

	if config.NumRounds == 0 {
		config.NumRounds = 200
	}

	if config.StartingStack == 0 {
		config.StartingStack = 1000
	}

	c.games[gameID] = newGame(config)
	return &proto.CreateGameResponse{GameId: gameID}, nil
}

func (c *Casino) JoinGame(ctx context.Context, request *proto.JoniGameRequest) (*proto.JoniGameResponse, error) {
	gameID := request.GameId

	game, ok := c.games[gameID]
	if !ok {
		return nil, fmt.Errorf("invalid gameID")
	}

	token, seatNumber := game.seatPlayer(request.PlayerId)
	return &proto.JoniGameResponse{
		Token:      fmt.Sprintf(tokenFormat, gameID, token),
		SeatNumber: uint32(seatNumber),
		PlayerId:   request.PlayerId,
	}, nil
}

func (c *Casino) ReceiveUpdates(ctx context.Context, request *proto.ReceiveUpdatesRequest) (*proto.ReceiveUpdatesResponse, error) {
	token := request.Token
	player, err := c.findPlayer(token)
	if err != nil {
		return nil, err
	}

	return player.FlushUpdates(), nil
}

func (c *Casino) Act(ctx context.Context, request *proto.ActRequest) (*proto.ActResponse, error) {
	player, err := c.findPlayer(request.Token)
	if err != nil {
		return nil, err
	}

	player.InputAction(request)
	return &proto.ActResponse{}, nil
}

func (c *Casino) findPlayer(token string) (*grpcBot, error) {
	if token == "" {
		return nil, fmt.Errorf("invalid token")
	}

	split := strings.Split(token, tokenGameDelimiter)
	if len(split) != 2 {
		return nil, fmt.Errorf("invalid token")
	}
	game, ok := c.games[split[0]]
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	player, ok := game.players[split[1]]
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return player, nil
}