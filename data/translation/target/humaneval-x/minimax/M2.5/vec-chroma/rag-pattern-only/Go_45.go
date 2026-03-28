package main

import "fmt"

func TriangleArea(a float64, h float64) float64 {
	return a * h / 2
}

func main() {
	// Example usage matching the Python doctest
	result := TriangleArea(5, 3)
	fmt.Println(result)
}
