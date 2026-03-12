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
		// Check for factors from 2 to sqrt(p)
		sqrtP := int(math.Sqrt(float64(p)))
		for k := 2; k <= sqrtP; k++ {
			if p%k == 0 {
				return false
			}
		}
		return true
	}

	f := []int{0, 1}
	for {
		// Calculate next Fibonacci number
		next := f[len(f)-1] + f[len(f)-2]
		f = append(f, next)

		// Check if it's prime, if so decrement counter
		if isPrime(next) {
			n--
		}
		if n == 0 {
			return next
		}
	}
}
