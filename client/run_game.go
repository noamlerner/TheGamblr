package client

import (
	"context"

	"github.com/noamlerner/TheGamblr/engine"
)

func RunLocalGameBetweenBots(playerBots []engine.BotPlayer, names []string) {
	ctx := context.Background()
	// We are going to be creating a GRPCBot that wraps around each of the BotPlayers. This GRPC bot knows how to
	// use a grpc connection to talk to the casino and feed information to the BotPlayer. It also knows how to use the
	// BotPlayer to take actions.
	grpcBots := make([]*GRPCBot, len(playerBots))
	for i, p := range playerBots {
		grpcBots[i] = NewClientBot(names[i], p, "localhost:8080")
	}

	// bot1 creates a new game.
	gameID := grpcBots[0].CreateGame(ctx)

	// All of our GRPCBots are going to join this game
	for _, p := range grpcBots {
		p.JoinGame(ctx, gameID)
	}

	// Bot1 starts the game
	grpcBots[0].StartGame(ctx)

	// async, we allow each bot to start running their game loop other than the one at index 0. We need one bot
	// to run its game loop in the main thread so that the program does end early. The rest need to be run async
	// or else they will be blocked by the first Bots game loop.
	for i := 1; i < len(grpcBots); i++ {
		grpcBot := grpcBots[i]
		go func() {
			grpcBot.Run(ctx)
		}()
	}

	// Bot at index 0 starts playing. We do not do this async so that the program doesn't end.
	grpcBots[0].Run(ctx)
}
