package analyzer

import (
	"errors"

	"golang.org/x/tools/go/analysis"
	"testcase.go/rules"
	//"golang.org/x/tools/go/analysis/passes/slog"
)

var Analyzer = &analysis.Analyzer{
	Name: "loganalyze",
	Doc:  "check logs for matching all rules",
	Run:  run,
}

type IncomeMessage struct {
	text string
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}

func CheckRules(pass *analysis.Pass, msg *IncomeMessage) bool {
	if !rules.IsStartsFromLowerCase(msg.text) {
		return false
	} else if !rules.IsEnglishLetter(msg.text) {
		return false
	} else if !rules.IsEmojiOrSpecialSymbol(msg.text) {
		return false
	} else if !rules.IsSensetiveData(msg.text) {
		return false
	}

	return true
}
