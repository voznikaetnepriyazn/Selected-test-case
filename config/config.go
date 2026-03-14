package config

import "golang.org/x/tools/go/analysis"

type Config struct {
	LowerLetterRule     bool
	IsEnglishRule       bool
	IsExtraSymbolsRule  bool
	IsSensetiveDataRule bool
}

func Init() *Config {
	return &Config{
		LowerLetterRule:     true,
		IsEnglishRule:       true,
		IsExtraSymbolsRule:  true,
		IsSensetiveDataRule: true,
	}
}

func Load(pass *analysis.Pass) *Config {
	cfg := Init()
	return cfg
}
