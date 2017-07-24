package luhn

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Valid tells if the given string is a valid Luhn code
func Valid(l string) bool {
	l = strings.Replace(l, " ", "", -1)
	if len(l) <= 1 {
		return false
	}

	sum := 0
	for i, r := range []rune(l) {
		if !unicode.IsDigit(r) {
			return false
		}

		d := int(r - '0')
		if i%2 == 1 || i == 0 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return sum%10 == 0
}
