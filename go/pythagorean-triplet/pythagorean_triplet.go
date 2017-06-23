// Package pythagorean provides operations on pythagorean triplets
package pythagorean

const testVersion = 1

// Triplet contains 3 numbers a,b,c
type Triplet struct {
	a, b, c int
}

// Range calculates all pythagorean triplets for a given range
func Range(min, max int) []Triplet {
	triplets := make([]Triplet, 0)

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if (a*a + b*b) == c*c {
					triplets = append(triplets, Triplet{a: a, b: b, c: c})
				}
			}
		}
	}
	return triplets
}

// Sum calculates all pythagorean triplets for which the 3 terms summed equal the given sum
func Sum(sum int) []Triplet {
	var triplets []Triplet
	for a := 1; a <= sum/3; a++ {
		for b := a; b <= (sum-a)/2; b++ {
			c := sum - a - b

			if (a*a + b*b) == c*c {
				triplets = append(triplets, Triplet{a: a, b: b, c: c})
			}
		}
	}
	return triplets
}
