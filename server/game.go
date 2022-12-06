package server

import (
	"github.com/noamlerner/TheGamblr/engine"

	"github.com/google/uuid"
)

type game struct {
	dealer            *engine.Dealer
	gameOver          bool
	existingPlayerIds map[string]bool
	players           map[string]*grpcBot
	id                string
}

func newGame(id string, config *engine.GameConfig) *game {
	dealer := engine.NewDealer(config)
	return &game{
		id:                id,
		dealer:            dealer,
		existingPlayerIds: map[string]bool{},
		players:           map[string]*grpcBot{},
	}
}

func (g *game) RuGame() {
	_ = g.dealer.RunGame()
	g.gameOver = true
}
func (g *game) seatPlayer(id string) (string, int) {
	token := uuid.Must(uuid.NewRandom()).String()
	grpcBot := newGrpcBot().(*grpcBot)
	seatNumber := g.dealer.SeatPlayer(id, grpcBot)
	g.players[token] = grpcBot
	return token, seatNumber
}
