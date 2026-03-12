package main

import (
	"fmt"
	"math"
)

// IsSimplePower checks if a number x is a simple power of n (n^k = x).
func IsSimplePower(x int, n int) bool {
	// Handle the edge case where n is 1. x must also be 1.
	if n == 1 {
		return x == 1
	}

	// n <= 0 (and not 1) can't produce a positive power for x > 1
	if n <= 0 {
		return false
	}

	// Special case for x=1, which is n^0
	if x == 1 {
		return true
	}

	val := float64(x)
	base := float64(n)

	// Calculate the logarithm of x to the base n.
	// log_n(x) = log_e(x) / log_e(n)
	p := math.Log(val) / math.Log(base)

	// If the result is very close to a whole number, check if n^round(p) == x.
	// Using a small epsilon to handle floating point inaccuracies.
	const epsilon = 1e-10
	if math.Abs(p-math.Round(p)) < epsilon {
		return int(math.Pow(base, math.Round(p))) == x
	}

	return false
}

// main function to test IsSimplePower
func main() {
	fmt.Printf("is_simple_power(1, 4) => %v\n", IsSimplePower(1, 4))       // Expected: true
	fmt.Printf("is_simple_power(2, 2) => %v\n", IsSimplePower(2, 2))       // Expected: true
	fmt.Printf("is_simple_power(8, 2) => %v\n", IsSimplePower(8, 2))       // Expected: true
	fmt.Printf("is_simple_power(3, 2) => %v\n", IsSimplePower(3, 2))       // Expected: false
	fmt.Printf("is_simple_power(3, 1) => %v\n", IsSimplePower(3, 1))       // Expected: false
	fmt.Printf("is_simple_power(5, 3) => %v\n", IsSimplePower(5, 3))       // Expected: false
	fmt.Printf("is_simple_power(27, 3) => %v\n", IsSimplePower(27, 3))      // Expected: true
	fmt.Printf("is_simple_power(625, 5) => %v\n", IsSimplePower(625, 5))    // Expected: true
	fmt.Printf("is_simple_power(624, 5) => %v\n", IsSimplePower(624, 5))    // Expected: false
}
