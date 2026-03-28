package main

import "fmt"

func IsSimplePower(x int, n int) bool {
	// Your task is to write a function that returns true if a number x is a simple
	// power of n and false in other cases.
	// x is a simple power of n if n**int=x
	// For example:
	// IsSimplePower(1, 4) => true
	// IsSimplePower(2, 2) => true
	// IsSimplePower(8, 2) => true
	// IsSimplePower(3, 2) => false
	// IsSimplePower(3, 1) => false
	// IsSimplePower(5, 3) => false
	if n == 1 {
		return x == 1
	}
	power := 1
	for power < x {
		// Check for overflow
		if power > x/n {
			return false
		}
		power = power * n
	}
	return power == x
}

func main() {
	// Test cases
	fmt.Println(IsSimplePower(1, 4)) // true
	fmt.Println(IsSimplePower(2, 2)) // true
	fmt.Println(IsSimplePower(8, 2)) // true
	fmt.Println(IsSimplePower(3, 2)) // false
	fmt.Println(IsSimplePower(3, 1)) // false
	fmt.Println(IsSimplePower(5, 3)) // false
}
