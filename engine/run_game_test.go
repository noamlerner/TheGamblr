package engine

import "testing"

func Test(t *testing.T) {
	config := NewDefaultGameConfig()
	config.LogLevel = LogLevelActions
	dealer := Dealer(config)
	dealer.SeatPlayer("Noam", &OneActionBot{action: CallAction})
	dealer.SeatPlayer("Bean", &OneActionBot{action: CallAction})
	dealer.SeatPlayer("TheGamblr", &OneActionBot{action: RaiseAction})
	dealer.RunGame()
}
