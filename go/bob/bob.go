// Package bob provides a teenager bot that has very limited answers
package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

// Hey is the means to interact with Bob, it will respond a string based on the input
func Hey(s string) string {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return "Fine. Be that way!"
	}

	if isAllUpper(s) {
		return "Whoa, chill out!"
	}

	if strings.HasSuffix(s, "?") {
		return "Sure."
	}

	return "Whatever."
}

func isAllUpper(s string) bool {
	hasLetters := false

	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			hasLetters = true
			if unicode.IsLower(r) {
				return false
			}
		}
	}
	return hasLetters
}
