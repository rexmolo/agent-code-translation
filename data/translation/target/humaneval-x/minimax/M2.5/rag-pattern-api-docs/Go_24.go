package main

import "fmt"

func LargestDivisor(n int) int {
	// Iterate from n-1 down to 1 to find the largest divisor
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			return i
		}
	}
	// If no divisor found (shouldn't happen for n > 1), return 1
	return 1
}

func main() {
	// Test the function
	fmt.Println(LargestDivisor(15)) // Expected output: 5
}
