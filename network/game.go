package network

import (
	"pokerengine/engine"

	"github.com/google/uuid"
)

type game struct {
	dealer *engine.Dealer

	existingPlayerIds map[string]bool
	players           map[string]*grpcBot
}

func newGame(config *engine.GameConfig) *game {
	dealer := engine.NewDealer(config)
	return &game{
		dealer:            dealer,
		existingPlayerIds: map[string]bool{},
	}
}

func (g *game) seatPlayer(id string) (string, int) {
	token := uuid.Must(uuid.NewRandom()).String()
	grpcBot := newGrpcBot()
	seatNumber := g.dealer.SeatPlayer(id, grpcBot)
	g.players[token] = grpcBot
	return token, seatNumber
}
