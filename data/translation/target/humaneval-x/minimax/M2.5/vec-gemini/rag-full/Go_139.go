package main

import "fmt"

func SpecialFactorial(n int) int {
	factI := 1
	specialFact := 1
	for i := 1; i <= n; i++ {
		factI *= i
		specialFact *= factI
	}
	return specialFact
}

func main() {
	// Test with example: special_factorial(4) = 1! * 2! * 3! * 4! = 1 * 2 * 6 * 24 = 288
	fmt.Println(SpecialFactorial(4))
}