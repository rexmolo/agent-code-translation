package main

import "fmt"

func LargestDivisor(n int) int {
	// Handle edge cases: if n is 0 or 1, there's no proper divisor smaller than n
	if n <= 1 {
		return 0
	}

	// Iterate from n-1 down to 1, find the largest divisor
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			return i
		}
	}

	// This should never be reached if n > 1
	return 0
}

func main() {
	// Example usage
	fmt.Println(LargestDivisor(15)) // Output: 5
}