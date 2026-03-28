package main

import (
	"math"
)

func Iscube(a int) bool {
	absA := a
	if absA < 0 {
		absA = -absA
	}

	// Calculate the cube root and round it
	cubeRoot := int64(math.Round(math.Pow(float64(absA), 1.0/3.0)))

	// Cube the rounded value and compare with the original
	return cubeRoot*cubeRoot*cubeRoot == int64(absA)
}