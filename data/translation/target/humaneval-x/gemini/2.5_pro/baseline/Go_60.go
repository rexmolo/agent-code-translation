package main

import "fmt"

// SumToN is a function that sums numbers from 1 to n.
func SumToN(n int) int {
	// The Python code `sum(range(n + 1))` sums integers from 0 to n.
	// This is equivalent to the formula for the sum of an arithmetic series.
	return n * (n + 1) / 2
}

// main function to demonstrate the usage of SumToN and match the docstring examples.
func main() {
	fmt.Println(SumToN(30))
	fmt.Println(SumToN(100))
	fmt.Println(SumToN(5))
	fmt.Println(SumToN(10))
	fmt.Println(SumToN(1))
}
