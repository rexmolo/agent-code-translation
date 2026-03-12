package main

import "fmt"

// Digits calculates the product of the odd digits of a positive integer n.
// It returns 0 if all digits are even.
// For example:
// Digits(1)  == 1
// Digits(4)  == 0
// Digits(235) == 15
func Digits(n int) int {
	product := 1
	hasOddDigit := false

	// The docstring specifies a positive integer, but we handle negative numbers by taking their absolute value.
	if n < 0 {
		n = -n
	}

	// If n is 0, the loop will not run and it will correctly return 0.
	tempN := n

	for tempN > 0 {
		digit := tempN % 10
		if digit%2 != 0 {
			product *= digit
			hasOddDigit = true
		}
		tempN /= 10
	}

	if !hasOddDigit {
		return 0
	}

	return product
}

// main function to run and test the Digits function with examples.
func main() {
	fmt.Println(Digits(1))
	fmt.Println(Digits(4))
	fmt.Println(Digits(235))
}
