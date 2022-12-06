package server

import (
	context "context"
	"fmt"
	"strings"

	"github.com/noamlerner/TheGamblr/engine"
	"github.com/noamlerner/TheGamblr/proto/thegamblr/proto"

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
		LogLevel:      engine.LogLevelCards,
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

	c.games[gameID] = newGame(gameID, config)
	return &proto.CreateGameResponse{GameId: gameID}, nil
}

func (c *Casino) JoinGame(ctx context.Context, request *proto.JoinGameRequest) (*proto.JoinGameResponse, error) {
	gameID := request.GameId

	game, ok := c.games[gameID]
	if !ok {
		return nil, fmt.Errorf("invalid gameID")
	}

	token, seatNumber := game.seatPlayer(request.PlayerId)
	return &proto.JoinGameResponse{
		Token:      fmt.Sprintf(tokenFormat, gameID, token),
		SeatNumber: uint32(seatNumber),
		PlayerId:   request.PlayerId,
	}, nil
}

func (c *Casino) StartGame(ctx context.Context, request *proto.StartGameRequest) (*proto.StartGameResponse, error) {
	game, _, err := c.findGameAndPlayer(request.Token)
	if err != nil {
		return nil, err
	}
	go func() {
		game.RuGame()
	}()
	return &proto.StartGameResponse{}, nil
}

func (c *Casino) ReceiveUpdates(ctx context.Context, request *proto.ReceiveUpdatesRequest) (*proto.ReceiveUpdatesResponse, error) {
	token := request.Token
	_, player, err := c.findGameAndPlayer(token)
	if err != nil {
		return nil, err
	}

	return player.FlushUpdates(int(request.SequenceNumber)), nil
}

func (c *Casino) Act(ctx context.Context, request *proto.ActRequest) (*proto.ActResponse, error) {
	game, player, err := c.findGameAndPlayer(request.Token)
	if err != nil {
		return nil, err
	}

	if game.gameOver {
		return nil, fmt.Errorf("game is over")
	}

	player.InputAction(request)
	return &proto.ActResponse{}, nil
}

func (c *Casino) findGameAndPlayer(token string) (*game, *grpcBot, error) {
	if token == "" {
		return nil, nil, fmt.Errorf("invalid token")
	}

	split := strings.Split(token, tokenGameDelimiter)
	if len(split) != 2 {
		return nil, nil, fmt.Errorf("invalid token")
	}
	game, ok := c.games[split[0]]
	if !ok {
		return nil, nil, fmt.Errorf("invalid token")
	}

	player, ok := game.players[split[1]]
	if !ok {
		return nil, nil, fmt.Errorf("invalid token")
	}
	return game, player, nil
}
