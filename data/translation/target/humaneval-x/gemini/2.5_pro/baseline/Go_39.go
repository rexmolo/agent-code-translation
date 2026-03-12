package main

import (
	"fmt"
)

// isPrime checks if a number is prime.
func isPrime(p int) bool {
	if p < 2 {
		return false
	}
	// A standard primality test is to check for factors up to the square root of p.
	// The condition k*k <= p is an efficient way to do this without floating point math.
	for k := 2; k*k <= p; k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}

// PrimeFib returns n-th number that is a Fibonacci number and it's also prime.
func PrimeFib(n int) int {
	a, b := 0, 1
	for {
		// Generate the next Fibonacci number
		a, b = b, a+b

		// Check if the new number is prime
		if isPrime(b) {
			n--
		}

		// If we have found the n-th prime Fibonacci, return it.
		if n == 0 {
			return b
		}
	}
}

// main function to run and test the PrimeFib function.
func main() {
	fmt.Println(PrimeFib(1))
	fmt.Println(PrimeFib(2))
	fmt.Println(PrimeFib(3))
	fmt.Println(PrimeFib(4))
	fmt.Println(PrimeFib(5))
}
