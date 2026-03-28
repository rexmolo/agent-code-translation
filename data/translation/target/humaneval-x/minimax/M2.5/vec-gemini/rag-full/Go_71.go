package main

import (
	"fmt"
	"math"
)

func TriangleArea(a float64, b float64, c float64) interface{} {
	// Check if three sides form a valid triangle
	// Sum of any two sides must be greater than the third side
	if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}

	// Calculate area using Heron's formula
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	// Round to 2 decimal places
	area = math.Floor(area*100+0.5) / 100

	return area
}

func main() {
	// Example usage
	result := TriangleArea(3, 4, 5)
	fmt.Println(result)

	result = TriangleArea(1, 2, 10)
	fmt.Println(result)
}