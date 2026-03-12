package main

import "fmt"

func LargestDivisor(n int) int {
	// Find the largest divisor of n that is smaller than n
	// Iterate from n-1 down to 1
	for i := n - 1; i > 0; i-- {
		if n%i == 0 {
			return i
		}
	}
	return 1 // 1 is always a divisor of any positive integer
}

func main() {
	// Example usage
	fmt.Println(LargestDivisor(15)) // Output: 5
}
