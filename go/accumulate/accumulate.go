package accumulate

const testVersion = 1

func Accumulate(strs []string, f func(string) string) []string {
	acc := make([]string, 0)

	for _, s := range strs {
		acc = append(acc, f(s))
	}
	return acc
}
