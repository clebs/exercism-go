package pangram

import "unicode"

const testVersion = 1

func IsPangram(s string) bool {
	alphabet := make(map[rune]struct{})
	runes := []rune(s)

	for _, r := range runes {
		if _, ok := alphabet[r]; !ok && unicode.IsLetter(r) && r < 128 {
			alphabet[r] = struct{}{}
		}
	}
	return len(alphabet) > 25
}
