package main

import "fmt"

// Fib returns the n-th Fibonacci number.
//
// Fibonacci sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	// Test cases from the Python docstrings
	fmt.Printf("fib(10) = %d\n", Fib(10)) // Expected: 55
	fmt.Printf("fib(1) = %d\n", Fib(1))   // Expected: 1
	fmt.Printf("fib(8) = %d\n", Fib(8))   // Expected: 21
}
