package main

import (
	"math"
)

func Iscube(a int) bool {
	a = int(math.Abs(float64(a)))
	// For small numbers, just iterate
	if a < 1000000000 {
		for i := 0; i*i*i <= a; i++ {
			if i*i*i == a {
				return true
			}
		}
		return false
	}
	// For large numbers, use floating-point with careful handling
	root := math.Round(math.Pow(float64(a), 1.0/3.0))
	// Check the rounded root and its neighbors for floating-point errors
	for i := int64(root) - 1; i <= int64(root)+1; i++ {
		if i*i*i == int64(a) {
			return true
		}
	}
	return false
}
