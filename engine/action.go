package engine

type Action int

const (
	FoldAction Action = iota
	CallAction
	RaiseAction
	CheckFoldAction
	SmallBlind
	BigBlind
)

var (
	actionVerb = map[Action]string{
		FoldAction:      "Folds",
		CallAction:      "Calls",
		RaiseAction:     "Raises To",
		CheckFoldAction: "Checks",
		SmallBlind:      "Pays Small Blind",
		BigBlind:        "Pays Big Blind",
	}
)

func (a Action) ActionVerb() string {
	return actionVerb[a]
}
