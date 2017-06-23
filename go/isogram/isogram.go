package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(s string) bool {
	found := make(map[rune]struct{}, 0)
	s = strings.ToLower(s)
	for _, r := range []rune(s) {
		if unicode.IsLetter(r) {
			if _, ok := found[r]; ok {
				return false
			}
			found[r] = struct{}{}
		}

	}
	return true
}
