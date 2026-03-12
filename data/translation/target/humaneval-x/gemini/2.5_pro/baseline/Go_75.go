package main

// isPrime is a helper function to check if a number is prime.
// It is an unexported function, a common Go pattern for helper functions.
// Note: The original Python is_prime implementation is incorrect for n < 2.
// This Go version corrects this, although it doesn't affect the behavior of
// IsMultiplyPrime, which only calls it for numbers >= 2.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for j := 2; j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

// IsMultiplyPrime returns true if the given number is the multiplication of 3 prime numbers
// and false otherwise. It's a Go translation of the provided Python function.
// The logic directly mirrors the Python version's triple-nested loop approach.
// Example: IsMultiplyPrime(30) == true (30 = 2 * 3 * 5)
func IsMultiplyPrime(a int) bool {
	// The Python loops `range(2,101)` iterate from 2 to 100 inclusive.
	// This is translated to `for ...; i <= 100; ...` in Go.
	for i := 2; i <= 100; i++ {
		if !isPrime(i) {
			continue
		}
		for j := 2; j <= 100; j++ {
			if !isPrime(j) {
				continue
			}
			for k := 2; k <= 100; k++ {
				if !isPrime(k) {
					continue
				}
				if i*j*k == a {
					return true
				}
			}
		}
	}
	return false
}
