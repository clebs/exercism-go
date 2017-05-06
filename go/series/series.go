package series

const testVersion = 2

// All returns all substrings of size n that can be made with the given string
func All(n int, s string) []string {
	r := []rune(s)
	all := make([]string, 0)
	for i := 0; i < len(r)-n+1; i++ {
		sx := string(r[i : i+n])
		all = append(all, sx)
	}
	return all
}

// UnsafeFirst returns the first substring of size n from the string s.
// This function does not check that n is not greater than the string and might panic.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of size n from the string s.
// This function checks that n is fits the given string and is safe to use with any input.
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
