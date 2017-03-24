package diffsquares

const testVersion = 1

// SumOfSquares returns the sum of the squares of the n first numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// SquareOfSums returns the square of the sum of the first n numbers
func SquareOfSums(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// Difference calculates the difference between the square of sums and the sum of squares of the first n natural numbers.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
