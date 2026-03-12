package main

import (
	"fmt"
	"math"
)

func TriangleArea(a float64, b float64, c float64) interface{} {
	// Check if triangle is valid: sum of any two sides must be greater than the third
	if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}

	// Heron's formula
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	// Round to 2 decimal places
	area = math.Round(area*100) / 100

	return area
}

func main() {
	fmt.Println(TriangleArea(3, 4, 5))
	fmt.Println(TriangleArea(1, 2, 10))
}
