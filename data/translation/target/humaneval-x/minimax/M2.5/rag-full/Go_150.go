package main

import "fmt"

func XOrY(n, x, y int) int {
	if n == 1 {
		return y
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return y
		}
	}
	return x
}

func main() {
	// Test cases
	fmt.Println(XOrY(7, 34, 12)) // Should return 34 (7 is prime)
	fmt.Println(XOrY(15, 8, 5))  // Should return 5 (15 is not prime)
}
