package main

import (
	"math"
)

func Iscube(a int) bool {
	if a < 0 {
		a = -a
	}
	cubeRoot := int(math.Round(math.Pow(float64(a), 1.0/3.0)))
	return cubeRoot*cubeRoot*cubeRoot == a
}
