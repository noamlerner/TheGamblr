package main

import (
	"github.com/noamlerner/TheGamblr/client"
	"github.com/noamlerner/TheGamblr/engine"
)

func main() {
	bot1 := engine.NewRandomActionBot()
	bot2 := engine.NewRandomActionBot()
	bot3 := engine.NewRandomActionBot()

	client.RunLocalGameBetweenBots([]engine.BotPlayer{bot1, bot2, bot3}, []string{"Noam", "Bean", "TheGamblr"})
}
