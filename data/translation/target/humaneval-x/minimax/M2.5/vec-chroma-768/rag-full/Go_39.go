package main

import (
	"math"
)

func PrimeFib(n int) int {
	// Helper function to check if a number is prime
	isPrime := func(p int) bool {
		if p < 2 {
			return false
		}
		// Check divisibility from 2 to sqrt(p)
		// Using int(math.Sqrt(float64(p))) + 1 to include the sqrt value in the range
		sqrtP := int(math.Sqrt(float64(p)))
		for k := 2; k <= sqrtP; k++ {
			if p%k == 0 {
				return false
			}
		}
		return true
	}

	// Start with Fibonacci sequence [0, 1]
	a, b := 0, 1

	// Generate Fibonacci numbers and count primes
	for {
		// Generate next Fibonacci number
		c := a + b
		a, b = b, c

		// Check if it's prime and decrement counter
		if isPrime(c) {
			n--
			if n == 0 {
				return c
			}
		}
	}
}