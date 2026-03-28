package main

import (
	"fmt"
	"math"
)

func PrimeFib(n int) int {
	// Helper function to check if a number is prime
	isPrime := func(p int) bool {
		if p < 2 {
			return false
		}
		// Check divisibility from 2 to min(sqrt(p)+1, p-1)
		limit := int(math.Sqrt(float64(p))) + 1
		if limit > p-1 {
			limit = p - 1
		}
		for k := 2; k <= limit; k++ {
			if p%k == 0 {
				return false
			}
		}
		return true
	}

	// Generate Fibonacci numbers iteratively
	a, b := 0, 1
	for {
		// Calculate next Fibonacci number
		a, b = b, a+b

		// Check if current Fibonacci number is prime
		if isPrime(a) {
			n--
		}
		if n == 0 {
			return a
		}
	}
}

func main() {
	// Test cases
	fmt.Println(PrimeFib(1)) // 2
	fmt.Println(PrimeFib(2)) // 3
	fmt.Println(PrimeFib(3)) // 5
	fmt.Println(PrimeFib(4)) // 13
	fmt.Println(PrimeFib(5)) // 89
}