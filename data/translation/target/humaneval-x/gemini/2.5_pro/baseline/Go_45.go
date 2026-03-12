package main

import "fmt"

// TriangleArea calculates the area of a triangle given the length of a side and its corresponding height.
func TriangleArea(a float64, h float64) float64 {
	return a * h / 2.0
}

func main() {
	// Example from the Python doctest
	result := TriangleArea(5, 3)
	fmt.Println(result) // Output: 7.5
}