package main

import (
	"fmt"
	"math"
)

// TruncateNumber takes a positive floating point number and can be decomposed into
// an integer part (largest integer smaller than the given number) and decimals
// (leftover part always smaller than 1).
// It returns the decimal part of the number.
func TruncateNumber(number float64) float64 {
	return math.Mod(number, 1.0)
}

// main function to demonstrate the usage of TruncateNumber
func main() {
	// Example from the Python docstring
	fmt.Println(TruncateNumber(3.5))
}
