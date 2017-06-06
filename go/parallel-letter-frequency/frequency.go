package letter

const testVersion = 1

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {
	partials := make(chan FreqMap, len(texts))
	for _, s := range texts {
		go func(s string) {
			partials <- Frequency(s)
		}(s)
	}

	result := FreqMap{}
	for range texts {
		for k, v := range <-partials {
			result[k] += v
		}
	}
	return result
}
