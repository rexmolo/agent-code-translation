package main

import (
	"math"
)

func Iscube(a int) bool {
	a = absInt(a)
	root := math.Round(math.Pow(float64(a), 1.0/3.0))
	return int(root)*int(root)*int(root) == a
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Test main to verify functionality
func main() {
	// Test cases from the examples
	testCases := []int{1, 2, -1, 64, 0, 180}
	for _, tc := range testCases {
		println(tc, Iscube(tc))
	}
}
