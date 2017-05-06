package queenattack

import (
	"fmt"
	"math"
)

const testVersion = 2

// CanQueenAttack tells if 2 chess queens can attack each other or not based on their position.
func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, fmt.Errorf("Both quens at the same position %s, only possible in quantum physics", w)
	}
	pw, err := toPoint(w)
	if err != nil {
		return false, err
	}
	pb, err := toPoint(b)
	if err != nil {
		return false, err
	}
	//same column or row
	if pw[0] == pb[0] || pw[1] == pb[1] {
		return true, nil
	}
	//diagonal
	if math.Abs(float64(pw[0]-pb[0])) == math.Abs(float64(pw[1]-pb[1])) {
		return true, nil

	}
	return false, nil
}

type point [2]int

func toPoint(s string) (point, error) {
	if len(s) != 2 {
		return point{}, fmt.Errorf("Incorrect chess piece position: %s", s)
	}
	var c point
	c[0] = int(s[0]) - 'a'
	c[1] = int(s[1]) - '1'

	if c[0] < 0 || c[0] > 7 || c[1] < 0 || c[1] > 7 {
		return point{}, fmt.Errorf("Queen out of board at %s", s)
	}
	return c, nil
}
