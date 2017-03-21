// Package triangle provides math operations related to triangle geometry
package triangle

import "math"

const testVersion = 3

// Kind is the enumeration of all distinct triangle types that can be recognized.
type Kind int

const (
	// NaT for Not a Triangle
	NaT Kind = iota
	// Equ for Equilateral: all sides are equally long
	Equ Kind = iota
	// Iso for Isosceles: 2 sides equal
	Iso Kind = iota
	// Sca for Scalene: All sides are different
	Sca Kind = iota
)

// KindFromSides tells if the 3 side lengths given form a triangle and what kind of triangle they form.
func KindFromSides(a, b, c float64) Kind {

	if notATriangle(a, b, c) {
		return NaT
	}

	if a == b || b == c || a == c {
		if a == b && a == c {
			return Equ
		}
		return Iso
	}
	return Sca
}

func notATriangle(a, b, c float64) bool {
	return (math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c)) || // side(s) not number(s)
		(math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1)) || // side(s) Inf
		(math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1)) || // side(s) -Inf
		(a <= 0 || b <= 0 || c <= 0) || (a > b+c || b > a+c || c > a+b) // triangle inequation not satisfied or no area
}
