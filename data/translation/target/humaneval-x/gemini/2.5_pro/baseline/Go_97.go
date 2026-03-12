package main

import "fmt"

// Multiply takes two integers and returns the product of their unit digits.
// The original Python code `abs(a % 10)` is translated by first taking
// the modulo and then ensuring the result is positive. This is necessary because
// Go's `%` operator can yield a negative result if the dividend is negative,
// unlike Python's `%` which takes the sign of the divisor.
func Multiply(a, b int) int {
	unitA := a % 10
	if unitA < 0 {
		unitA = -unitA
	}

	unitB := b % 10
	if unitB < 0 {
		unitB = -unitB
	}

	return unitA * unitB
}

// main function to demonstrate the Multiply function
func main() {
	fmt.Println(Multiply(148, 412))   // Expected: 16
	fmt.Println(Multiply(19, 28))     // Expected: 72
	fmt.Println(Multiply(2020, 1851)) // Expected: 0
	fmt.Println(Multiply(14, -15))    // Expected: 20
}
