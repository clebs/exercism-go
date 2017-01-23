//Package hamming calculates mutation by comparing DNA strands
package hamming

import "errors"

const testVersion = 5

// Distance calculates the distance between the DNA strands a and b by counting the number of different nucleotide
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA strands are not compatible, they have different lengths")
	}

	mutations := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			mutations++
		}
	}
	return mutations, nil
}
