package main

import (
	"context"

	"github.com/noamlerner/TheGamblr/client"
	"github.com/noamlerner/TheGamblr/engine"
)

func main() {
	// ClientBot takes an algorithm that implements engine.BotPlayer and handles the logic of talking to the casino.
	// Here, we are creating to ClientBots that take random actions.
	bot1 := client.NewClientBot("Bean", engine.NewRandomActionBot(), "localhost:8080")
	bot2 := client.NewClientBot("Noam", engine.NewRandomActionBot(), "localhost:8080")

	ctx := context.Background()
	// bot1 creates a new game and then both bot1 and bot2 join the game.
	gameID := bot1.CreateGame(ctx)
	bot1.JoinGame(ctx, gameID)
	bot2.JoinGame(ctx, gameID)

	// Bot1 starts the game
	bot1.StartGame(ctx)

	// async, bot1 starts playing. We do this async because there is a while loop in here, and bot2 also needs to start
	// playing.
	go func() {
		bot1.Run(ctx)
	}()

	// bot2 starts playing. We do not do this async so that the program doesn't end.
	bot2.Run(ctx)
}
