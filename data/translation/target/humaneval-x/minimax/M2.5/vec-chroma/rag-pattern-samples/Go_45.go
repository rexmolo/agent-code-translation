package main

import "fmt"

func TriangleArea(a float64, h float64) float64 {
	return a * h / 2.0
}

func main() {
	// Test case from docstring: triangle_area(5, 3) = 7.5
	fmt.Println(TriangleArea(5, 3))
}