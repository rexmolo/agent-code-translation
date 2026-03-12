package main

import "fmt"

// Fib4 computes the n-th element of the Fib4 number sequence.
// The sequence is defined as follows:
// fib4(0) -> 0
// fib4(1) -> 0
// fib4(2) -> 2
// fib4(3) -> 0
// fib4(n) -> fib4(n-1) + fib4(n-2) + fib4(n-3) + fib4(n-4).
// This is an efficient, non-recursive implementation.
func Fib4(n int) int {
	if n < 0 {
		// Assuming non-negative input as per sequence definition.
		// Returning 0 for negative n is a reasonable default.
		return 0
	}

	results := []int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	a, b, c, d := results[0], results[1], results[2], results[3]

	for i := 4; i <= n; i++ {
		a, b, c, d = b, c, d, a+b+c+d
	}

	return d
}

func main() {
	// Test cases from the original Python docstring
	fmt.Println(Fib4(5))
	fmt.Println(Fib4(6))
	fmt.Println(Fib4(7))
}
