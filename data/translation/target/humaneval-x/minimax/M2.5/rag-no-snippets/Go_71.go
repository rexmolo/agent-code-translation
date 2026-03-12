package main

import (
	"fmt"
	"math"
)

func TriangleArea(a float64, b float64, c float64) interface{} {
	// Three sides make a valid triangle when the sum of any two sides is greater than the third side
	if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	// Round to 2 decimal places
	area = math.Floor(area*100+0.5) / 100
	return area
}

func main() {
	fmt.Println(TriangleArea(3, 4, 5))
	fmt.Println(TriangleArea(1, 2, 10))
}
