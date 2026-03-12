package main

import "fmt"

func LargestPrimeFactor(n int) int {
	// Helper function to check if k is prime
	isPrime := func(k int) bool {
		if k < 2 {
			return false
		}
		// Check divisibility from 2 up to k-2 (equivalent to Python's range(2, k-1))
		for i := 2; i < k-1; i++ {
			if k%i == 0 {
				return false
			}
		}
		return true
	}

	largest := 1
	for j := 2; j <= n; j++ {
		if n%j == 0 && isPrime(j) {
			if j > largest {
				largest = j
			}
		}
	}
	return largest
}

func main() {
	fmt.Println(LargestPrimeFactor(13195)) // Output: 29
	fmt.Println(LargestPrimeFactor(2048))  // Output: 2
}
