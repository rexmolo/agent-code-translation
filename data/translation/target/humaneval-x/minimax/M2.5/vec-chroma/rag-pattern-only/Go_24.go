package main

import "fmt"

func LargestDivisor(n int) int {
	// Iterate from n-1 down to 1 to find the largest divisor
	for i := n - 1; i >= 1; i-- {
		if n%i == 0 {
			return i
		}
	}
	return 0 // No divisor found (edge case for n <= 1)
}

func main() {
	// Example usage
	fmt.Println(LargestDivisor(15)) // Output: 5
}
