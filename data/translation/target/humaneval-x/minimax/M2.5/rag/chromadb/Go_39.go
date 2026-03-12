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
		// Check divisibility from 2 to sqrt(p)
		limit := int(math.Sqrt(float64(p))) + 1
		for k := 2; k < limit && k < p; k++ {
			if p%k == 0 {
				return false
			}
		}
		return true
	}

	f := []int{0, 1}
	for {
		// Compute next Fibonacci number
		next := f[len(f)-1] + f[len(f)-2]
		f = append(f, next)

		// Check if it's prime and decrement counter
		if isPrime(next) {
			n--
		}

		// Return when we've found the nth prime Fibonacci
		if n == 0 {
			return next
		}
	}
}

func main() {
	fmt.Println(PrimeFib(1)) // 2
	fmt.Println(PrimeFib(2)) // 3
	fmt.Println(PrimeFib(3)) // 5
	fmt.Println(PrimeFib(4)) // 13
	fmt.Println(PrimeFib(5)) // 89
}
