package main

import (
	"math"
)

func TriangleArea(a float64, b float64, c float64) interface{} {
	// Check if the three sides form a valid triangle
	// Triangle inequality: sum of any two sides must be greater than the third
	if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}

	// Calculate semi-perimeter
	s := (a + b + c) / 2

	// Calculate area using Heron's formula
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	// Round to 2 decimal places
	roundedArea := math.Round(area*100) / 100

	return roundedArea
}