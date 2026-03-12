package main

import "fmt"

// LargestDivisor finds the largest integer that divides n evenly, smaller than n.
func LargestDivisor(n int) int {
	// The original Python code has a bug for n <= 1, where it attempts
	// division by zero. This Go implementation is more robust.
	// We loop from n-1 down to 1.
	for i := n - 1; i > 0; i-- {
		if n%i == 0 {
			// The first divisor found will be the largest because we are iterating downwards.
			return i
		}
	}
	// This part is only reached if n <= 1, as any integer n > 1 has 1 as a divisor.
	// For prime numbers, the loop will run until i = 1, and 1 will be returned.
	// For n <= 1, the problem is not well-defined. Returning 1 is a reasonable default.
	return 1
}

func main() {
	// Example from the Python docstring
	fmt.Println(LargestDivisor(15))
}
