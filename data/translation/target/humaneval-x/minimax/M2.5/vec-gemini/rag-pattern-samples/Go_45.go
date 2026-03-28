package main

import "fmt"

func TriangleArea(a float64, h float64) float64 {
	return a * h / 2.0
}

func main() {
	result := TriangleArea(5, 3)
	fmt.Println(result)
}
