package grains

import (
	"errors"
	"math"
)

const testVersion = 1

// Square returns the number of grains in the given table square by calculating 2^(n-1)
func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("Given square is out of the board")
	}

	return 1 << uint64(n-1), nil
}

// Total returns the max number of grains in a chessboard, aka 2^64
func Total() uint64 {
	return math.MaxUint64
}
