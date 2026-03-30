package main

import (
	"fmt"
	"math"
)

func TruncateNumber(number float64) float64 {
	return math.Mod(number, 1.0)
}

func main() {
	// Test the function
	result := TruncateNumber(3.5)
	fmt.Println(result)
}