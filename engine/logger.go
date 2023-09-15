package engine

import "fmt"

const (
	actionLogStatement                   = "%s %s %v\n"                               // PlayerID Action Amount
	boardLogStatement                    = "%s: \n\tCommunityCards: %s, \n\tPot %v\n" // Stage: CommunityCards: Cards, Pot:amount
	chipCountLogStatement                = "Stack Sizes: "
	playerChipCountLogStatement          = "\t%s - %v\n"
	winnersLogStatement                  = "%s won %v chips with %s making a %s\n"
	LogLevelNone                LogLevel = iota
	LogLevelWinners
	LogLevelStages
	LogLevelActions
)

type logger struct {
	logLevel LogLevel
}

func (l *logger) Action(a VisibleAction) {
	if l.logLevel < LogLevelActions {
		return
	}

	fmt.Printf(actionLogStatement, a.Player().Id(), a.ActionTaken().ActionVerb(), a.Amount())
}

func (l *logger) Stage(board ActiveBoard) {
	if l.logLevel < LogLevelStages {
		return
	}
	fmt.Printf(boardLogStatement, board.Stage().Name(), board.CommunityCards().String(), board.Pot())
}

func (l *logger) Winners(board RoundResults) {
	if l.logLevel < LogLevelWinners {
		return
	}
	for _, p := range board.PlayerResults() {
		if len(p.Cards()) > 0 {
			fmt.Printf(winnersLogStatement, p.Id(), p.ChipsWon(), p.Cards().String(), p.HandStrength().String())
		}
	}

	fmt.Println(chipCountLogStatement)
	for _, p := range board.PlayerResults() {
		fmt.Printf(playerChipCountLogStatement, p.Id(), p.Stack())
	}
}