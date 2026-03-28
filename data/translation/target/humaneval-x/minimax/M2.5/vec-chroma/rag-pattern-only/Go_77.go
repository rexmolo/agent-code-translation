package main

import (
	"math"
)

func Iscube(a int) bool {
	a = abs(a)
	cubeRoot := math.Cbrt(float64(a))
	rounded := int(math.Round(cubeRoot))
	return rounded*rounded*rounded == a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
