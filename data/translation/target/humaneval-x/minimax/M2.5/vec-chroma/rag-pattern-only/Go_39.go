package main

import (
	"fmt"
	"math"
)

func PrimeFib(n int) int {
	if n <= 0 {
		return 0
	}

	isPrime := func(p int) bool {
		if p < 2 {
			return false
		}
		if p == 2 {
			return true
		}
		if p%2 == 0 {
			return false
		}
		// Only check odd divisors up to sqrt(p)
		sqrtP := int(math.Sqrt(float64(p)))
		for k := 3; k <= sqrtP; k += 2 {
			if p%k == 0 {
				return false
			}
		}
		return true
	}

	f := []int{0, 1}
	count := 0

	for {
		// Generate next Fibonacci number
		next := f[len(f)-1] + f[len(f)-2]
		f = append(f, next)

		if isPrime(next) {
			count++
			if count == n {
				return next
			}
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