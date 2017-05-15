package etl

import "strings"

const testVersion = 1

// Transform converts old scrabble score representations into the new format
func Transform(m map[int][]string) map[string]int {
	r := make(map[string]int, 0)
	for k, v := range m {
		for _, l := range v {
			l = strings.ToLower(l)
			if _, ok := r[l]; !ok {
				r[l] = k
			}
		}
	}
	return r
}
