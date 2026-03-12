package main

import "fmt"

// TriangleArea calculates the area of a triangle given the base (a) and height (h).
// Formula: area = (base * height) / 2
//
// Example:
// TriangleArea(5, 3) returns 7.5
func TriangleArea(a float64, h float64) float64 {
	return a * h / 2.0
}

func main() {
	result := TriangleArea(5, 3)
	fmt.Printf("%.1f\n", result)
}
