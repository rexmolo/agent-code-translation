package main

import (
	"math"
)

func Iscube(a int) bool {
	a = int(math.Abs(float64(a)))
	root := int(math.Round(math.Cbrt(float64(a))))
	return root*root*root == a
}
