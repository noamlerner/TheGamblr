package engine

import "fmt"

const (
	roundNumberLogStatement              = "Round %v\n"
	twoStringFormatStatement             = "%s %s\n"
	cardsLogStatement                    = "%s received %s\n"
	actionLogStatement                   = "%s %s %v\n"                               // PlayerID ActionType Amount
	boardLogStatement                    = "%s: \n\tCommunityCards: %s, \n\tPot %v\n" // Stage: CommunityCards: Cards, Pot:amount
	chipCountLogStatement                = "Stack Sizes: "
	playerChipCountLogStatement          = "\t%s - %v\n"
	winnersLogStatement                  = "%s won %v chips with %s making a %s\n"
	LogLevelNone                LogLevel = iota
	LogLevelWinners
	LogLevelStages
	LogLevelActions
	LogLevelCards
)

type logger struct {
	logLevel LogLevel
}

func (l *logger) Cards(p *playerState) {
	if l.logLevel < LogLevelCards {
		return
	}
	fmt.Printf(cardsLogStatement, p.id, p.cards.String())
}
func (l *logger) Action(a Action) {
	if l.logLevel < LogLevelActions {
		return
	}

	if a.Type() == CheckFoldAction || a.Type() == FoldAction {
		fmt.Printf(twoStringFormatStatement, a.Player().Id(), a.Type().ActionVerb())
		return
	}

	fmt.Printf(actionLogStatement, a.Player().Id(), a.Type().ActionVerb(), a.Amount())
}

func (l *logger) Round(r int) {
	if l.logLevel < LogLevelStages {
		return
	}
	fmt.Printf(roundNumberLogStatement, r+1)
}
func (l *logger) Stage(board BoardState) {
	if l.logLevel < LogLevelStages {
		return
	}
	fmt.Printf(boardLogStatement, board.Stage().Name(), board.CommunityCards().String(), board.Pot())
}

func (l *logger) Winners(board BoardState) {
	if l.logLevel < LogLevelWinners {
		return
	}
	for _, p := range board.Players() {
		if p == nil {
			continue
		}
		if len(p.PlayerRoundResults().Cards()) > 0 {
			fmt.Printf(winnersLogStatement, p.Id(), p.PlayerRoundResults().ChipsWon(), p.PlayerRoundResults().Cards().String(), p.PlayerRoundResults().HandStrength().String())
		}
	}

	fmt.Println(chipCountLogStatement)
	for _, p := range board.Players() {
		if p == nil {
			continue
		}
		fmt.Printf(playerChipCountLogStatement, p.Id(), p.Stack())
	}
}
