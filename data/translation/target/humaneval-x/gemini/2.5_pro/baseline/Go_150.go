package main

import (
	"fmt"
)

// XOrY returns the value of x if n is a prime number and should return the value of y otherwise.
func XOrY(n, x, y int) int {
	// In primality tests, 1 is not considered a prime number.
	if n <= 1 {
		return y
	}

	// Check for factors from 2 up to n-1.
	// A more efficient check would go up to sqrt(n), but we are matching the original logic.
	for i := 2; i < n; i++ {
		if n%i == 0 {
			// If a factor is found, n is not prime.
			return y
		}
	}

	// If the loop completes without finding any factors, n is prime.
	return x
}

func main() {
	// Example 1 from the docstring
	fmt.Println(XOrY(7, 34, 12)) // Expected: 34

	// Example 2 from the docstring
	fmt.Println(XOrY(15, 8, 5)) // Expected: 5

	// Additional test cases
	fmt.Println(XOrY(2, 100, 200)) // Expected: 100
	fmt.Println(XOrY(1, 10, 20))   // Expected: 20
}
