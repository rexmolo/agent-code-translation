package main

import (
	"fmt"
	"math"
)

func TriangleArea(a float64, b float64, c float64) interface{} {
	// Check if the three sides form a valid triangle
	if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}

	// Calculate semi-perimeter using Heron's formula
	s := (a + b + c) / 2

	// Calculate area
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	// Round to 2 decimal places
	area = math.Round(area*100) / 100

	return area
}

func main() {
	// Test cases
	fmt.Println(TriangleArea(3, 4, 5))   // Expected: 6
	fmt.Println(TriangleArea(1, 2, 10)) // Expected: -1
}
