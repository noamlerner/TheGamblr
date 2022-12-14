package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type pokerTableTest struct {
	name           string
	playersAtSeats []int
}

var twoPlayerTest = &pokerTableTest{
	"Two PlayerResults",
	[]int{3, 6},
}
var tests = []*pokerTableTest{
	{
		"Full Table",
		[]int{0, 1, 2, 3, 4, 5, 6, 7},
	},
	{
		"Three PlayerResults",
		[]int{3, 4, 7},
	},
	{
		"Four PlayerResults",
		[]int{0, 3, 5, 7},
	},
	{
		"Five PlayerResults",
		[]int{3, 4, 5, 6, 7},
	},
	{
		"Six PlayerResults",
		[]int{2, 3, 4, 5, 6, 7},
	},
	{
		"Seven PlayerResults",
		[]int{0, 2, 3, 4, 5, 6, 7},
	},
}

func TestDealer_NewRound(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)

			dealer := DealerWithDefaultConfig()
			dealer.board.players = players
			dealer.newRound()

			smallBlindIndex := 0
			if test.playersAtSeats[0] == 0 {
				smallBlindIndex = 1
			}

			assert.Equal(t, test.playersAtSeats[smallBlindIndex], dealer.board.smallBlindButton)
			assert.Equal(t, 95, dealer.board.playerAtSeat(test.playersAtSeats[smallBlindIndex]).stack)
			assert.Equal(t, 5, dealer.board.playerAtSeat(test.playersAtSeats[smallBlindIndex]).chipsEnteredThisRound)
			assert.Equal(t, 90, dealer.board.playerAtSeat(test.playersAtSeats[smallBlindIndex+1]).stack)
			assert.Equal(t, 10, dealer.board.playerAtSeat(test.playersAtSeats[smallBlindIndex+1]).chipsEnteredThisRound)
			assert.Equal(t, 15, dealer.board.pot)

			smallBlindIndex++
			bigBlindIndex := smallBlindIndex + 1
			if bigBlindIndex >= len(test.playersAtSeats) {
				bigBlindIndex = 0
			}
			dealer.board.playerAtSeat(test.playersAtSeats[smallBlindIndex]).stack = 2
			dealer.newRound()
			assert.Equal(t, test.playersAtSeats[smallBlindIndex], dealer.board.smallBlindButton)
			assert.Equal(t, 0, dealer.board.playerAtSeat(test.playersAtSeats[smallBlindIndex]).stack)
			assert.Equal(t, 2, dealer.board.playerAtSeat(test.playersAtSeats[smallBlindIndex]).chipsEnteredThisStage)
			assert.Equal(t, 90, dealer.board.playerAtSeat(test.playersAtSeats[bigBlindIndex]).stack)
			assert.Equal(t, 10, dealer.board.playerAtSeat(test.playersAtSeats[bigBlindIndex]).chipsEnteredThisRound)
			assert.Equal(t, 12, dealer.board.pot)
		})
	}
}

func TestBoard_NewRound_TwoPlayers(t *testing.T) {
	twoPlayerIndexes := twoPlayerTest.playersAtSeats
	players, _ := setupPlayers(twoPlayerIndexes)

	dealer := DealerWithDefaultConfig()
	dealer.board.players = players
	dealer.newRound()

	assert.Equal(t, twoPlayerIndexes[0], dealer.board.smallBlindButton)
	assert.Equal(t, 95, dealer.board.playerAtSeat(twoPlayerIndexes[0]).stack)
	assert.Equal(t, 90, dealer.board.playerAtSeat(twoPlayerIndexes[1]).stack)
	assert.Equal(t, 15, dealer.board.pot)

	dealer.board.playerAtSeat(twoPlayerIndexes[1]).stack = 2
	dealer.newRound()
	assert.Equal(t, twoPlayerIndexes[1], dealer.board.smallBlindButton)
	assert.Equal(t, 0, dealer.board.playerAtSeat(twoPlayerIndexes[1]).stack)
	assert.Equal(t, 85, dealer.board.playerAtSeat(twoPlayerIndexes[0]).stack)
	assert.Equal(t, 12, dealer.board.pot)
}

func TestDealer_BettingRound_EveryoneCalls(t *testing.T) {
	for _, test := range append(tests, twoPlayerTest) {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)

			dealer := DealerWithDefaultConfig()
			dealer.board.players = players
			dealer.newRound()

			dealer.betting()
			assert.Equal(t, len(test.playersAtSeats)*dealer.gameConfig.SmallBlind*2, dealer.board.pot)

			for _, p := range test.playersAtSeats {
				assert.Equal(t, dealer.board.playerAtSeat(p).chipsEnteredThisStage, 10)
				assert.Equal(t, dealer.board.playerAtSeat(p).chipsEnteredThisRound, 10)
				assert.Equal(t, dealer.board.playerAtSeat(p).stack, 90)
			}

			dealer.board.nextStage()
			dealer.betting()
			assert.Equal(t, len(test.playersAtSeats)*dealer.gameConfig.SmallBlind*2, dealer.board.pot)
			for _, p := range test.playersAtSeats {
				assert.Equal(t, dealer.board.playerAtSeat(p).chipsEnteredThisStage, 0)
				assert.Equal(t, dealer.board.playerAtSeat(p).chipsEnteredThisRound, 10)
				assert.Equal(t, dealer.board.playerAtSeat(p).stack, 90)
			}
		})
	}
}

func TestDealer_BettingRound_OneRaiser(t *testing.T) {
	for _, test := range append(tests, twoPlayerTest) {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)
			dealer := DealerWithDefaultConfig()
			dealer.log.logLevel = LogLevelCards
			dealer.board.players = players
			raiserIndex := dealer.board.nextActiveSeat(0)
			// He raises and min raise will be 10, so everything has to get to 20.
			dealer.board.playerAtSeat(raiserIndex).actor = &OneActionBot{action: RaiseAction}
			dealer.newRound()

			dealer.betting()
			assert.Equal(t, len(test.playersAtSeats)*20, dealer.board.pot)
			for _, i := range test.playersAtSeats {
				p := dealer.board.playerAtSeat(i)
				if i == raiserIndex || i == dealer.board.nextActiveSeat(raiserIndex) {
					// since 1 raised, it only gets to go once. 1 is small blind, 2 is big blind.
					// This means 2's first call is to 1's raise, so it also only goes once.
					assert.Equal(t, 1, p.actor.(*OneActionBot).numCalled)
				} else {
					assert.Equal(t, 2, p.actor.(*OneActionBot).numCalled)
				}
				assert.Equal(t, 20, p.chipsEnteredThisStage)
				assert.Equal(t, 80, p.stack)
			}

			dealer.board.nextStage()
			dealer.betting()
			assert.Equal(t, len(test.playersAtSeats)*30, dealer.board.pot)
			for _, p := range test.playersAtSeats {
				assert.Equal(t, dealer.board.playerAtSeat(p).chipsEnteredThisStage, 10)
				assert.Equal(t, dealer.board.playerAtSeat(p).chipsEnteredThisRound, 30)
				assert.Equal(t, dealer.board.playerAtSeat(p).stack, 70)
			}
		})
	}
}

func TestDealer_BettingRound_RaiseToInfinity(t *testing.T) {
	for _, test := range append(tests, twoPlayerTest) {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)
			dealer := DealerWithDefaultConfig()
			dealer.board.players = players

			raiserIndex := dealer.board.nextActiveSeat(0)
			dealer.board.playerAtSeat(raiserIndex).actor = &OneActionBot{action: RaiseAction}
			raiserIndex = dealer.board.nextActiveSeat(raiserIndex)
			dealer.board.playerAtSeat(raiserIndex).actor = &OneActionBot{action: RaiseAction}

			dealer.newRound()
			dealer.betting()
			assert.Equal(t, len(test.playersAtSeats)*100, dealer.board.pot)
			for _, i := range test.playersAtSeats {
				p := dealer.board.playerAtSeat(i)
				assert.Equal(t, 100, p.chipsEnteredThisStage)
				assert.Equal(t, 0, p.stack)
			}

			dealer.board.nextStage()
			dealer.betting()
			assert.Equal(t, len(test.playersAtSeats)*100, dealer.board.pot)
			for _, p := range test.playersAtSeats {
				assert.Equal(t, dealer.board.playerAtSeat(p).chipsEnteredThisStage, 0)
				assert.Equal(t, dealer.board.playerAtSeat(p).chipsEnteredThisRound, 100)
				assert.Equal(t, dealer.board.playerAtSeat(p).stack, 0)
			}
		})
	}
}

func TestDealer_BettingRound_Folder(t *testing.T) {
	for _, test := range append(tests, twoPlayerTest) {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)
			dealer := DealerWithDefaultConfig()
			dealer.board.players = players

			folder := dealer.board.playerAtSeat(dealer.board.nextActiveSeat(0))
			folder.actor = &OneActionBot{action: FoldAction}
			raiser := dealer.board.playerAtSeat(dealer.board.nextActiveSeat(folder.seatNumber))
			raiser.actor = &OneActionBot{action: RaiseAction}

			dealer.newRound()
			dealer.betting()

			if len(test.playersAtSeats) == 2 {
				return
			}
			assert.Equal(t, raiser.actor.(*OneActionBot).numCalled, folder.actor.(*OneActionBot).numCalled)
			// small blind
			assert.Equal(t, folder.chipsEnteredThisStage, 5)
			assert.Equal(t, folder.chipsEnteredThisRound, 5)
			assert.Equal(t, folder.status, PlayerStatusFolded)

			if len(test.playersAtSeats) == 2 {
				return
			}

			folder = dealer.board.playerAtSeat(dealer.board.prevActiveSeat(raiser.seatNumber))
			folder.actor.(*OneActionBot).action = FoldAction

			dealer.board.nextStage()
			dealer.betting()

			assert.Equal(t, 3, folder.actor.(*OneActionBot).numCalled)
			assert.Equal(t, folder.chipsEnteredThisStage, 0)
			assert.Equal(t, folder.chipsEnteredThisRound, 20)
			assert.Equal(t, folder.status, PlayerStatusFolded)
		})
	}
}

func TestDealer_BettingRound_AllFolded(t *testing.T) {
	for _, test := range append(tests, twoPlayerTest) {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)
			dealer := DealerWithDefaultConfig()
			dealer.board.players = players

			for _, i := range test.playersAtSeats {
				dealer.board.playerAtSeat(i).actor.(*OneActionBot).action = FoldAction
			}

			dealer.newRound()
			dealer.board.nextStage()
			dealer.betting()

			numSeen := 0
			dealer.board.iterateActivePlayers(func(p *playerState) {
				numSeen++
				if numSeen == len(test.playersAtSeats) {
					// we short circuit when theres only one player left
					assert.Equal(t, p.actor.(*OneActionBot).numCalled, 0)
				} else {
					assert.Equal(t, p.actor.(*OneActionBot).numCalled, 1)
				}
			})
		})
	}
}

func TestDealer_PlayRound_NoOneFolds(t *testing.T) {
	for _, test := range append(tests, twoPlayerTest) {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)
			dealer := DealerWithDefaultConfig()
			dealer.board.players = players
			winners := dealer.playRound()

			assert.Equal(t, RoundOver, dealer.board.stage)
			assert.Len(t, dealer.board.communityCards, 5)
			assert.NotNil(t, winners)

			for i := 0; i < winners.Len()-1; i++ {
				// we should always be better or the same as the next hand
				assert.True(t, winners[i].hand.Beats(winners[i+1].hand) || winners[i].hand.Tie(winners[i+1].hand))
				assert.Len(t, winners[i].p.cards, 2)
			}
		})
	}
}

func TestDealer_PlayRound_EveryoneFolds(t *testing.T) {
	for _, test := range append(tests, twoPlayerTest) {
		t.Run(test.name, func(t *testing.T) {
			players, _ := setupPlayers(test.playersAtSeats)
			dealer := DealerWithDefaultConfig()
			dealer.board.players = players
			dealer.board.iterateActivePlayers(func(p *playerState) {
				p.actor.(*OneActionBot).action = FoldAction
			})

			winners := dealer.playRound()

			assert.Equal(t, RoundOver, dealer.board.stage)
			assert.Len(t, dealer.board.communityCards, 0)
			assert.NotNil(t, winners)
			assert.Len(t, winners, 1)
		})
	}
}

func TestDealer_CashOutRound_EveryoneElseFolded(t *testing.T) {
	dealer := DealerWithDefaultConfig()
	dealer.board.pot = 100
	winners := []*playerForWinnerCalculations{{
		p: &playerState{
			visiblePlayerState: visiblePlayerState{
				stack:  100,
				status: PlayerStatusPlaying,
			},
		},
	}}
	dealer.cashOutRound(winners)

	assert.Equal(t, 200, winners[0].p.stack)
}

func TestDealer_CashOutRound(t *testing.T) {
	var tests = []struct {
		name        string
		winners     []*playerForWinnerCalculations
		stacks      []int
		carrOverPot int
	}{
		{
			name: "WinningHand",
			winners: []*playerForWinnerCalculations{
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack: 100,

							status: PlayerStatusPlaying},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateFlush(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							status: PlayerStatusPlaying,
						},

						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateHighCard(Ace)),
				},
			},
			stacks: []int{220, 100, 100},
		},
		{
			name: "WinningHand_AllIn",
			winners: []*playerForWinnerCalculations{
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  0,
							id:     "1",
							status: PlayerStatusAllIn,
						},
						chipsEnteredThisRound: 20,
					},
					hand: NewHand(GenerateFlush(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "2",
							status: PlayerStatusPlaying,
							stack:  100,
						},
						chipsEnteredThisRound: 40},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "3",
							stack:  100,
							status: PlayerStatusPlaying,
						},

						chipsEnteredThisRound: 10,
					},
					hand: NewHand(GenerateHighCard(Ace)),
				},

				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "4",
							stack:  100,
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateHighCard(King)),
				},
			},
			stacks: []int{70, 140, 100, 100},
		},
		{
			name: "SplitPot_TwoWay",
			winners: []*playerForWinnerCalculations{
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "1",
							stack:  100,
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "2",
							stack:  100,
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							id:     "3",
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 10,
					},
					hand: NewHand(GenerateHighCard(Ace)),
				},

				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							id:     "4",
							status: PlayerStatusPlaying,
						},

						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateHighCard(King)),
				},
			},
			stacks: []int{165, 165, 100, 100},
		},
		{
			name: "SplitPot_ThreeWay",
			winners: []*playerForWinnerCalculations{
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:    "1",
							stack: 100,

							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "2",
							stack:  100,
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							id:     "3",
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							id:     "4",
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 40,
					},
					hand: NewHand(GenerateHighCard(King)),
				},
			},
			stacks:      []int{153, 153, 153, 100},
			carrOverPot: 1,
		},
		{
			name: "SplitPot_Threeway_OneAllIn",
			winners: []*playerForWinnerCalculations{
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "1",
							stack:  0,
							status: PlayerStatusAllIn,
						},
						chipsEnteredThisRound: 30},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "2",
							stack:  100,
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 60},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack: 100,
							id:    "3",

							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 60,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},

				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							id:     "4",
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 60,
					},
					hand: NewHand(GenerateHighCard(King)),
				},
			},
			stacks:      []int{40, 185, 185, 100},
			carrOverPot: 0,
		}, {
			name: "SplitPot_Then_Split_Pot",
			winners: []*playerForWinnerCalculations{
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "1",
							stack:  0,
							status: PlayerStatusAllIn,
						},
						chipsEnteredThisRound: 30,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							id:     "2",
							stack:  0,
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 30,
					},
					hand: NewHand(GenerateStraightTo(Ace)),
				},
				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							id:     "3",
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 60,
					},
					hand: NewHand(GenerateStraightTo(Six)),
				},

				{
					p: &playerState{
						visiblePlayerState: visiblePlayerState{
							stack:  100,
							id:     "4",
							status: PlayerStatusPlaying,
						},
						chipsEnteredThisRound: 60,
					},
					hand: NewHand(GenerateHighCard(Six)),
				},
			},
			stacks:      []int{60, 60, 130, 130},
			carrOverPot: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dealer := DealerWithDefaultConfig()
			players := make([]*playerState, len(test.winners))
			for i, w := range test.winners {
				dealer.board.pot += w.p.chipsEnteredThisRound
				players[i] = w.p
			}
			dealer.board.players = players
			dealer.cashOutRound(test.winners)

			for i, w := range test.winners {
				assert.Equal(t, test.stacks[i], w.p.stack)
			}
			assert.Equal(t, test.carrOverPot, dealer.carryOverPot)
		})
	}
}
