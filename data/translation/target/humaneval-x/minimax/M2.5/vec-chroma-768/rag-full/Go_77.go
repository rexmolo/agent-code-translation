package main

import "math"

func Iscube(a int) bool {
	absA := int(math.Abs(float64(a)))
	root := int(math.Round(math.Pow(float64(absA), 1.0/3.0)))
	return root*root*root == absA
}