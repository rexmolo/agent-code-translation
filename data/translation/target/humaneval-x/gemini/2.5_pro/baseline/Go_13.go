package main

import (
	"fmt"
)

// GreatestCommonDivisor returns the greatest common divisor of two integers a and b
// using the Euclidean algorithm.
func GreatestCommonDivisor(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// main is the entry point of the program.
// It demonstrates the usage of the GreatestCommonDivisor function with examples.
func main() {
	// Example 1 from the Python docstring
	fmt.Println(GreatestCommonDivisor(3, 5))

	// Example 2 from the Python docstring
	fmt.Println(GreatestCommonDivisor(25, 15))
}
