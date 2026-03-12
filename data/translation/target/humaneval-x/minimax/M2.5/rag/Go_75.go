package main

import "fmt"

func IsMultiplyPrime(a int) bool {
	// Helper function to check if a number is prime
	isPrime := func(n int) bool {
		for j := 2; j < n; j++ {
			if n%j == 0 {
				return false
			}
		}
		return true
	}

	for i := 2; i <= 100; i++ {
		if !isPrime(i) {
			continue
		}
		for j := 2; j <= 100; j++ {
			if !isPrime(j) {
				continue
			}
			// Early termination: if i*j already exceeds a, no need to check k
			if i*j > a {
				break
			}
			for k := 2; k <= 100; k++ {
				if !isPrime(k) {
					continue
				}
				if i*j*k == a {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	// Test the function
	fmt.Println(IsMultiplyPrime(30)) // Should print: true (30 = 2 * 3 * 5)
}
