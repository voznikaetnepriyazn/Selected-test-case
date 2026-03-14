package rules

import "unicode"

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
