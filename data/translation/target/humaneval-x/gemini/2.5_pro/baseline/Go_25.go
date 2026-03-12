package main

import (
	"fmt"
)

// Factorize returns a slice of prime factors of a given integer in ascending order.
// Each factor is listed as many times as it appears in the factorization.
func Factorize(n int) []int {
	// In Python: fact = []
	factors := []int{}

	// In Python: i = 2
	i := 2

	// The Python code uses `while i <= int(math.sqrt(n) + 1):`.
	// A more idiomatic and performant Go equivalent is `for i*i <= n`.
	// This avoids floating-point math in the loop condition and is correct because
	// if a number n has a factor larger than its square root, it must also have a factor smaller than it.
	// The loop correctly handles the changing value of n.
	for i*i <= n {
		// In Python: if n % i == 0:
		if n%i == 0 {
			// In Python: fact.append(i)
			factors = append(factors, i)
			// In Python: n //= i
			n /= i
		} else {
			// In Python: i += 1
			i++
		}
	}

	// This handles the case where the remaining n is a prime number greater than 1.
	// In Python: if n > 1: fact.append(n)
	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

// main function to demonstrate the Factorize function, similar to Python's doctests.
func main() {
	fmt.Println(Factorize(8))
	fmt.Println(Factorize(25))
	fmt.Println(Factorize(70))
}
