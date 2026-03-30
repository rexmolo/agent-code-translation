package main

import "fmt"

func main() {
	// Test cases
	fmt.Println(XOrY(7, 34, 12)) // Expected: 34
	fmt.Println(XOrY(15, 8, 5)) // Expected: 5
}

func XOrY(n, x, y int) int {
	// Return y if n is 1 (1 is not prime by convention)
	if n == 1 {
		return y
	}
	// Check if n is divisible by any number from 2 to n-1
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return y // n is not prime
		}
	}
	// If no divisor found, n is prime, return x
	return x
}
