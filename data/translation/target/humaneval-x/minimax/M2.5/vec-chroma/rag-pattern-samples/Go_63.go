package main

import (
	"fmt"
)

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

	// Use iterative approach for efficiency (O(n) instead of exponential)
	// Keep track of the last three values
	a, b, c := 0, 0, 1 // fibfib(0), fibfib(1), fibfib(2)
	for i := 3; i <= n; i++ {
		next := a + b + c
		a, b, c = b, c, next
	}
	return c
}

func main() {
	// Test cases
	fmt.Println(Fibfib(1)) // 0
	fmt.Println(Fibfib(5)) // 4
	fmt.Println(Fibfib(8)) // 24
}