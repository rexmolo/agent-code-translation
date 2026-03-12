package main

import "fmt"

// isPrime checks if a number is prime.
// This is a direct translation of the helper function in the Python code.
// The original Python loop `range(2, k - 1)` checks for divisors from 2 up to k-2.
// This implementation preserves that exact, though unconventional, logic.
func isPrime(k int) bool {
	if k < 2 {
		return false
	}
	// For k=2 or k=3, the loop condition `i < k-1` is false, so the loop
	// doesn't run, and the function correctly returns true.
	for i := 2; i < k-1; i++ {
		if k%i == 0 {
			return false
		}
	}
	return true
}

// LargestPrimeFactor returns the largest prime factor of n.
// This function's logic is a direct translation of the provided Python code.
// It follows the docstring's assumption that n > 1 and is not a prime.
func LargestPrimeFactor(n int) int {
	largest := 1
	// The loop iterates through all possible factors from 2 up to n.
	for j := 2; j <= n; j++ {
		// If j is a factor of n and j is prime, it's a prime factor.
		if n%j == 0 && isPrime(j) {
			// Since j is always increasing, any new prime factor found will be the largest so far.
			// This simplifies Python's `largest = max(largest, j)` to `largest = j`.
			largest = j
		}
	}
	return largest
}

// main function to demonstrate the usage of LargestPrimeFactor
// and replicate the examples from the Python docstring.
func main() {
	fmt.Println(LargestPrimeFactor(13195))
	fmt.Println(LargestPrimeFactor(2048))
}
