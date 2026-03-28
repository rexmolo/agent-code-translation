package main

import "fmt"

func TruncateNumber(number float64) float64 {
	return number % 1
}

func main() {
	// Test case from docstring
	fmt.Println(TruncateNumber(3.5)) // Expected: 0.5
}