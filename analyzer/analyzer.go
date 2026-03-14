package analyzer

import (
	"errors"

	"golang.org/x/tools/go/analysis"
	//"golang.org/x/tools/go/analysis/passes/slog"
)

var Analyzer = &analysis.Analyzer{
	Name: "loganalyze",
	Doc:  "check logs for matching all rules",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}
