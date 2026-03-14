package analyzer

import (
	//"errors"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"testcase.go/rules"

	"golang.org/x/tools/go/ast/inspector"
	//"golang.org/x/tools/go/analysis/passes/slog"
)

var Analyzer = &analysis.Analyzer{
	Name:     "loganalyze",
	Doc:      "check logs for matching all rules",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

type IncomeMessage struct {
	text string
}

func run(pass *analysis.Pass) (interface{}, error) {
	//instead tree inspect call
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	//filter only funcs
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		be := n.(*ast.CallExpr)
		if !IsIncomeMessageLog(be, pass) { //need to add func that define income message like the log
			return
		}

		//need to add func that extract message from log
		messages := ExtractMessagesFromLog(be, pass)

		for _, message := range messages {
			CheckRules(pass, message)
		}
	})

	return nil, nil
}

func IsIncomeMessageLog(be *ast.CallExpr, pass *analysis.Pass) bool {

}

func ExtractMessagesFromLog(be *ast.CallExpr, pass *analysis.Pass) []*IncomeMessage {

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
