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
	// Test with n=4, expected output: 288
	fmt.Println(SpecialFactorial(4))
}