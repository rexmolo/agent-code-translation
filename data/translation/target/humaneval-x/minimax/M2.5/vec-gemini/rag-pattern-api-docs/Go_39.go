package main

import (
	"math"
)

func PrimeFib(n int) int {
	// isPrime checks if a number is prime
	isPrime := func(p int) bool {
		if p < 2 {
			return false
		}
		// Check divisibility from 2 to sqrt(p)
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

	// Generate Fibonacci numbers
	// Using two variables instead of a slice for efficiency
	a, b := 0, 1
	for {
		// Calculate next Fibonacci number
		a, b = b, a+b
		// Check if it's prime
		if isPrime(b) {
			n--
			if n == 0 {
				return b
			}
		}
	}
}