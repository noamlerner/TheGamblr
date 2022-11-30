package engine

type ActionType int

const (
	FoldAction ActionType = iota
	CallAction
	RaiseAction
	CheckFoldAction
	SmallBlind
	BigBlind
)

var (
	actionVerb = map[ActionType]string{
		FoldAction:      "Folds",
		CallAction:      "Calls",
		RaiseAction:     "Raises To",
		CheckFoldAction: "Checks",
		SmallBlind:      "Pays Small Blind",
		BigBlind:        "Pays Big Blind",
	}
)

func (a ActionType) ActionVerb() string {
	return actionVerb[a]
}
