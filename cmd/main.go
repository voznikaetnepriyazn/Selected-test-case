package main

import (
	_ "testcase/analyzer"

	"golang.org/x/tools/go/analysis/passes/findcall"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(findcall.Analyzer)
}
