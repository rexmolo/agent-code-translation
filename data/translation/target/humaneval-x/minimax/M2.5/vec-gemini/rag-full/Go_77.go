package main

import "math"

func Iscube(a int) bool {
	a = int(math.Abs(float64(a)))
	cubeRoot := math.Pow(float64(a), 1.0/3.0)
	rounded := int(math.Round(cubeRoot))
	return rounded*rounded*rounded == a
}