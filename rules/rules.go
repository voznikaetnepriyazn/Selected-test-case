package rules

import (
	"strings"
	"unicode"
)

func IsStartsFromLowerCase(s string) bool {
	if len(s) == 0 {
		return true
	}

	for _, r := range s {
		if unicode.IsSpace(r) || unicode.IsDigit(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
			continue
		}
		return unicode.IsLower(r)
	}

	return true
}

func IsEnglishOnly(s string) bool {
	if len(s) == 0 {
		return true
	}

	for _, r := range s {
		if r > 0x7F {
			return false
		}
	}

	return true
}

func IsEmojiOrSpecialSymbol(s string) bool {
	for _, i := range s {
		if unicode.IsPunct(i) || unicode.IsSymbol(i) {
			return false
		}
	}

	return true
}

var SensetiveData = []string{
	"api_key",
	"token",
	"password",
}

func IsSensetiveData(s string) bool {
	s = strings.ToLower(s)

	for _, i := range SensetiveData {
		if strings.Contains(s, i) {
			return false
		}
	}

	return true
}
