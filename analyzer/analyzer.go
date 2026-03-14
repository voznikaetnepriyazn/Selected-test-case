package analyzer

import (
	//"errors"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"testcase.go/config"
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
	text            string
	pos             token.Pos
	IsConcatenation bool
}

func run(pass *analysis.Pass) (interface{}, error) {
	//instead tree inspect call
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	cfg := config.Load(pass)

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
			CheckRules(pass, cfg, message)
		}
	})

	return nil, nil
}

func IsIncomeMessageLog(be *ast.CallExpr, pass *analysis.Pass) bool {

}

func ExtractMessagesFromLog(be *ast.CallExpr, pass *analysis.Pass) []*IncomeMessage {

}

func CheckRules(pass *analysis.Pass, cfg *config.Config, msg *IncomeMessage) {
	if config.Init().LowerLetterRule && !rules.IsStartsFromLowerCase(msg.text) {
		pass.Reportf(msg.pos, "log message must starts from lower case")
	} else if config.Init().IsEnglishRule && !rules.IsEnglishLetter(msg.text) {
		pass.Reportf(msg.pos, "log message must contains only english letters")
	} else if config.Init().IsExtraSymbolsRule && !rules.IsEmojiOrSpecialSymbol(msg.text) {
		pass.Reportf(msg.pos, "log message must not contains emoji or punctuation symbols")
	} else if config.Init().IsSensetiveDataRule && !rules.IsSensetiveData(msg.text) {
		pass.Reportf(msg.pos, "log message must not contains sensetive data")
	}
}
