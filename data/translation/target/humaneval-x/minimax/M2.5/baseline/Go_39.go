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
		// Note: Python's min(int(sqrt(p)) + 1, p-1) simplifies to just sqrt(p)+1
		// since we only need to check up to sqrt(p) for primality
		limit := int(math.Sqrt(float64(p))) + 1
		for k := 2; k < limit; k++ {
			if p%k == 0 {
				return false
			}
		}
		return true
	}

	f := []int{0, 1}
	for {
		// Generate next Fibonacci number
		next := f[len(f)-1] + f[len(f)-2]
		f = append(f, next)
		if isPrime(next) {
			n--
		}
		if n == 0 {
			return next
		}
	}
}

func main() {
	// Test the function
	fmt.Println(PrimeFib(1)) // 2
	fmt.Println(PrimeFib(2)) // 3
	fmt.Println(PrimeFib(3)) // 5
	fmt.Println(PrimeFib(4)) // 13
	fmt.Println(PrimeFib(5)) // 89
}
