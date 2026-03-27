package main

import "fmt"

func TriangleArea(a float64, h float64) float64 {
	return a * h / 2.0
}

func main() {
	fmt.Println(TriangleArea(5, 3))
}
