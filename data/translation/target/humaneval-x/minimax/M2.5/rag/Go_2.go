package main

import (
	"math"
)

func TruncateNumber(number float64) float64 {
	return math.Mod(number, 1.0)
}

func main() {
	// Example usage
	result := TruncateNumber(3.5)
	println(result) // Output: 0.5
}
