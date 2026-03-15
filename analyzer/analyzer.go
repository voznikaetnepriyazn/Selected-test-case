package analyzer

import (
	//"errors"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"testcase/config"
	"testcase/rules"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"

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
	Text            string
	Pos             token.Pos
	IsConcatenation bool
}

var SupportedLoggers = map[string][]string{
	"log":             {"Print", "Printf", "Fatal", "Fatalf", "Panic", "Panicf", "Printf"},
	"log/slog":        {"Info", "Warn", "Debug", "Error", "Log", "LogAttrs"},
	"go.uber.org/zap": {"Info", "Debug", "Warn", "Error", "Panic", "Fatal"},
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
		if !IsIncomeMessageLog(be, pass) {
			return
		}

		messages := ExtractMessagesFromLog(be, pass)

		for _, message := range messages {
			CheckRules(pass, cfg, message)
		}
	})

	return nil, nil
}

func IsIncomeMessageLog(be *ast.CallExpr, pass *analysis.Pass) bool {
	sel, ok := be.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	methodName := sel.Sel.Name

	typeInfo := pass.TypesInfo.Types[sel.X]

	if !typeInfo.IsValue() {
		return false
	}

	if pkgName, ok := sel.X.(*ast.Ident); ok {
		if obj := pass.TypesInfo.Uses[pkgName]; obj != nil {
			if pkg, ok := obj.(*types.PkgName); ok {
				pkgPath := pkg.Imported().Path()

				if methods, exists := SupportedLoggers[pkgPath]; exists {
					for _, m := range methods {
						if m == methodName {
							return true
						}
					}
				}
			}
		}
	}

	if typ := typeInfo.Type; typ != nil {
		if ptr, ok := typ.(*types.Pointer); ok {
			typ = ptr.Elem()
		}

		if named, ok := typ.(*types.Named); ok {
			typeObj := named.Obj()
			if typeObj == nil {
				return false
			}

			pkg := typeObj.Pkg()
			typeName := typeObj.Name()

			if pkg != nil {
				pkgPath := pkg.Path()
				if methods, exists := SupportedLoggers[pkgPath]; exists {
					for _, m := range methods {
						if m == methodName {
							return true
						}
					}
				}
			}

			_ = typeName
		}
	}

	return false
}

func ExtractMessagesFromLog(be *ast.CallExpr, pass *analysis.Pass) []*IncomeMessage {
	var messages []*IncomeMessage

	for i, arg := range be.Args {
		if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			text := strings.Trim(lit.Value, `"`)

			messages = append(messages, &IncomeMessage{
				Text:            text,
				Pos:             lit.Pos(),
				IsConcatenation: false,
			})
			continue
		}

		if bin, ok := arg.(*ast.BinaryExpr); ok && bin.Op == token.ADD {
			if lit, ok := bin.X.(*ast.BasicLit); ok && lit.Kind == token.STRING {
				text := strings.Trim(lit.Value, `"`)

				messages = append(messages, &IncomeMessage{
					Text:            text,
					Pos:             lit.Pos(),
					IsConcatenation: true,
				})
			}
		}
		_ = i
	}

	return messages
}

func CheckRules(pass *analysis.Pass, cfg *config.Config, msg *IncomeMessage) {
	if cfg.LowerLetterRule && !rules.IsStartsFromLowerCase(msg.Text) {
		pass.Reportf(msg.Pos, "log message must starts from lower case")
	} else if cfg.IsEnglishRule && !rules.IsEnglishOnly(msg.Text) {
		pass.Reportf(msg.Pos, "log message must contains only english letters")
	} else if cfg.IsExtraSymbolsRule && !rules.IsEmojiOrSpecialSymbol(msg.Text) {
		pass.Reportf(msg.Pos, "log message must not contains emoji or punctuation symbols")
	} else if cfg.IsSensetiveDataRule && !rules.IsSensetiveData(msg.Text) {
		pass.Reportf(msg.Pos, "log message must not contains sensetive data")
	}
}
