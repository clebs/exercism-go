//Package raindrops provides functions to get information from rain drops
package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 2

const (
	pling = "Pling"
	plang = "Plang"
	plong = "Plong"
)

// Convert n drops into a string representing the sound they will make when they hit
func Convert(n int) string {
	var buffer bytes.Buffer

	if n%3 == 0 {
		buffer.WriteString(pling)
	}
	if n%5 == 0 {
		buffer.WriteString(plang)
	}
	if n%7 == 0 {
		buffer.WriteString(plong)
	}

	if buffer.Len() == 0 {
		return strconv.Itoa(n)
	}

	return buffer.String()
}
