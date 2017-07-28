package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

const testVersion = 2

var normalizer = regexp.MustCompile(`\W+`)

// Encode takes an arbitrary string and encodes it by the square code method.
func Encode(s string) string {
	s = normalizer.ReplaceAllString(s, "")
	s = strings.ToLower(s)
	l := len(s)

	if l < 2 {
		return s
	}

	cf := math.Sqrt(float64(l))
	c := int(math.Ceil(cf))

	square := make([]string, c)

	for i := 0; i < l; i++ {
		square[i%c] += s[i : i+1]
	}

	return strings.Join(square, " ")
}
