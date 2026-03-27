package main

import (
	"fmt"
)

func TruncateNumber(number float64) float64 {
	return number % 1.0
}

func main() {
	// Test the function
	fmt.Println(TruncateNumber(3.5)) // Should output 0.5
}