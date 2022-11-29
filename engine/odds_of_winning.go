package engine

import "strconv"

// OddsOfWinningOrTie takes an array of cards, each of len 2. This array can be of len 1 or greater.
// We run n simulation to calculate the odds of winning.
// You can set as many community cards as you want between 0 and 5.
// We will pass back a float for the probability of winning for each passed in hand.
// So for example, if you have 5 players at a table and you are wondering what the probability of someone with AA vs
// someone with QQ is of winning, you can pass in
// []Cards{{AA}, {QQ} }, []Cards {}, 5, 1000
// and we will return []float{}
func OddsOfWinningOrTie(hands []Cards, numPlayers int, communityCards Cards, n int) []float64 {
	orderedDeck := NewHardOrderedDeck(hands, numPlayers, communityCards)
	d := DealerWithDefaultConfig()
	d.gameConfig.StartingStack = 100000000

	for i := 0; i < numPlayers; i++ {
		d.SeatPlayer(strconv.Itoa(i), &OneActionBot{action: CallAction})
	}

	d.deck = orderedDeck

	winCount := make([]float64, len(hands))
	for i := 0; i < n; i++ {
		winners := d.playRound()
		for _, w := range winners {
			if w.chipsWon == 0 {
				continue
			}

			seatsFromSmallBlind := w.p.seatNumber - d.board.smallBlindButton
			if seatsFromSmallBlind < 0 {
				seatsFromSmallBlind += numPlayers
			}
			if seatsFromSmallBlind < len(winCount) {
				winCount[seatsFromSmallBlind]++
			}
		}
	}

	for i, w := range winCount {
		winCount[i] = w / float64(n)
	}
	return winCount
}
