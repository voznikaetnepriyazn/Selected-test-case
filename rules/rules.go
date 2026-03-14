package rules

import (
	"strings"
	"unicode"
)

func IsStartsFromLowerCase(s string) bool {
	if unicode.IsLetter(rune(s[0])) && unicode.IsLower(rune(s[0])) {
		return true
	}

	return false
}

func IsEnglishLetter(s string) bool {
	for _, i := range s {
		if unicode.IsLetter(i) || unicode.IsDigit(i) || unicode.IsSpace(i) {
			return true
		}
	}

	return false
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
