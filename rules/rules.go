package rules

import "unicode"

func IsStartsFromLowerCase(s string) bool {
	for _, i := range s {
		if unicode.IsLetter(i) {
			return unicode.IsLower(i)
		}
		if !unicode.IsSpace(i) {
			return true
		}
	}
	return true
}
