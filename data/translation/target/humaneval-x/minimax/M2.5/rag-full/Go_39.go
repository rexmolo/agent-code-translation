package main

import (
	"math"
)

func PrimeFib(n int) int {
	// isPrime helper function to check if a number is prime
	isPrime := func(p int) bool {
		if p < 2 {
			return false
		}
		// Check divisibility from 2 to sqrt(p)
		limit := int(math.Sqrt(float64(p)))
		for k := 2; k <= limit; k++ {
			if p%k == 0 {
				return false
			}
		}
		return true
	}

	// Fibonacci sequence: a=0, b=1
	a, b := 0, 1

	for {
		// Calculate next Fibonacci number
		a, b = b, a+b

		// Check if it's prime
		if isPrime(a) {
			n--
			if n == 0 {
				return a
			}
		}
	}
}