package main

import "fmt"

func TriangleArea(a float64, h float64) float64 {
	return a * h / 2.0
}

func main() {
	// Test with the example from the docstring: triangle_area(5, 3) = 7.5
	result := TriangleArea(5, 3)
	fmt.Printf("TriangleArea(5, 3) = %.1f\n", result)
}