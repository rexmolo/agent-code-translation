package main

import "fmt"

// Fibfib computes the n-th element of the FibFib sequence
// FibFib is defined as:
// fibfib(0) == 0
// fibfib(1) == 0
// fibfib(2) == 1
// fibfib(n) == fibfib(n-1) + fibfib(n-2) + fibfib(n-3)
func Fibfib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return Fibfib(n-1) + Fibfib(n-2) + Fibfib(n-3)
}

func main() {
	// Test cases
	fmt.Println(Fibfib(1)) // Expected: 0
	fmt.Println(Fibfib(5)) // Expected: 4
	fmt.Println(Fibfib(8)) // Expected: 24
}
