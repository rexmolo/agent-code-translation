package main

import "fmt"

// SpecialFactorial calculates the Brazilian factorial.
// The Brazilian factorial is defined as:
// brazilian_factorial(n) = n! * (n-1)! * (n-2)! * ... * 1!
// where n > 0
//
// For example:
// SpecialFactorial(4) -> 288
//
// The function receives an integer as input and returns the special
// factorial of this integer.
func SpecialFactorial(n int) int {
	factI := 1
	specialFact := 1
	for i := 1; i <= n; i++ {
		factI *= i
		specialFact *= factI
	}
	return specialFact
}

// main function to demonstrate the usage of SpecialFactorial
func main() {
	// Example from the Python docstring
	n := 4
	result := SpecialFactorial(n)
	fmt.Printf("Special factorial of %d is %d\n", n, result)
}
