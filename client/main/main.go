package main

import (
	"context"

	"pokerengine/client"
	"pokerengine/engine"
)

func main() {
	bot1 := client.NewClientBot("Bean", engine.NewRandomActionBot(), "localhost:8080")
	bot2 := client.NewClientBot("Noam", engine.NewRandomActionBot(), "localhost:8080")
	ctx := context.Background()
	gameID := bot1.CreateGame(ctx)
	bot1.JoinGame(ctx, gameID)
	bot2.JoinGame(ctx, gameID)

	bot1.StartGame(ctx)
	go func() {
		bot1.Run(ctx)
	}()
	bot2.Run(ctx)
}
