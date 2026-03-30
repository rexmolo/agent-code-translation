package main

// PrimeLength returns true if the length of the string is a prime number,
// false otherwise.
func PrimeLength(s string) bool {
	l := len(s)
	if l == 0 || l == 1 {
		return false
	}
	for i := 2; i < l; i++ {
		if l%i == 0 {
			return false
		}
	}
	return true
}
