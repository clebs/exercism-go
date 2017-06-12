package summultiples

const testVersion = 2

// SumMultiples calculates the sum of all numbers that are multiple of any of the given divisors up to the given limit.
func SumMultiples(limit int, divisors ...int) int {
	partials := make(chan []int, len(divisors))
	for _, d := range divisors {
		go func(divisor int) {
			ml := make([]int, 0)
			for i := 1; i*divisor < limit; i++ {
				ml = append(ml, i*divisor)
			}
			partials <- ml
		}(d)
	}

	total := 0
	added := make(map[int]bool, 0)
	for range divisors {
		for _, p := range <-partials {
			if !added[p] {
				added[p] = true
				total += p
			}
		}
	}
	return total
}
