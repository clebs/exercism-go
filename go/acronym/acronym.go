//Package acronym provides an API to generate acronyms from sentences or names.
package acronym

import (
	"bytes"
	"unicode"
)

const testVersion = 2
const wordsRegex = "\\w+"

// Abbreviate turns a String with several words into an acronym.
func Abbreviate(s string) string {
	var acronym bytes.Buffer

	// iteration is done this way to support unicode characters that are more than 1 byte long
	var prev rune
	for _, l := range s {
		if prev == 0 {
			acronym.WriteRune(unicode.ToUpper(l))
			prev = l
			continue
		}

		if (unicode.IsLetter(l) && !unicode.IsLetter(prev)) || (unicode.IsUpper(l) && !unicode.IsUpper(prev)) {
			acronym.WriteRune(unicode.ToUpper(l))
		}
		prev = l
	}
	return acronym.String()
}
