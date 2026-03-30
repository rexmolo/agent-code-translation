package main

import (
	"math"
)

func TriangleArea(a float64, b float64, c float64) interface{} {
	if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}
	s := (a + b + c) / 2
	area := s * (s - a) * (s - b) * (s - c)
	sqrtArea := math.Sqrt(area)
	roundedArea := math.Round(sqrtArea*100) / 100
	return roundedArea
}
