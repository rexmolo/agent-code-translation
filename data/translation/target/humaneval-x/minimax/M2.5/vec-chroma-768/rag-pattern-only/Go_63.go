package main

import "fmt"

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

	a, b, c := 0, 0, 1 // represents fibfib(n-3), fibfib(n-2), fibfib(n-1)
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

func main() {
	// Test the function
	testCases := []int{0, 1, 2, 5, 8}
	for _, n := range testCases {
		fmt.Printf("fibfib(%d) = %d\n", n, Fibfib(n))
	}
}
