package main

import "fmt"

func main() {
	// Example usage
	fmt.Println(XOrY(7, 34, 12)) // Output: 34
	fmt.Println(XOrY(15, 8, 5)) // Output: 5
}

func XOrY(n, x, y int) int {
	// Returns x if n is prime, y otherwise
	// Note: 1 is not considered prime, so return y
	if n == 1 {
		return y
	}
	// Check if n is divisible by any number from 2 to n-1
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return y // n is not prime
		}
	}
	return x // n is prime
}
