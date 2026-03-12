package main

import "math"

// TruncateNumber returns the decimal part of a positive floating point number.
func TruncateNumber(number float64) float64 {
	return math.Mod(number, 1.0)
}

func main() {
	// Example usage
	result := TruncateNumber(3.5)
	println(result)
}
