package helpers

// Remove delete an element from the provided array.
func Remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Equal check if provided arrays are equal.
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
