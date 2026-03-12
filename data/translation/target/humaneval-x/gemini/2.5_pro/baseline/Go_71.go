package main

import (
	"fmt"
	"math"
)

// TriangleArea calculates the area of a triangle given its three sides.
// It returns the area rounded to 2 decimal points if the sides form a valid triangle.
// Otherwise, it returns -1.
// A triangle is valid if the sum of any two sides is greater than the third side,
// and all sides have a positive length.
func TriangleArea(a float64, b float64, c float64) interface{} {
	// Check for triangle inequality and positive side lengths
	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		return -1
	}

	// Calculate the semi-perimeter
	s := (a + b + c) / 2

	// Calculate the area using Heron's formula
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	// Round the area to 2 decimal places
	roundedArea := math.Round(area*100) / 100

	return roundedArea
}

// main function to demonstrate the usage of TriangleArea
func main() {
	// Example 1: Valid triangle
	result1 := TriangleArea(3, 4, 5)
	fmt.Printf("Area of triangle with sides 3, 4, 5: %.2f\n", result1)

	// Example 2: Invalid triangle
	result2 := TriangleArea(1, 2, 10)
	fmt.Printf("Area of triangle with sides 1, 2, 10: %v\n", result2)
}
