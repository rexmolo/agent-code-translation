package main

import (
	"fmt"
)

// Monotonic returns true if a slice of integers is monotonically increasing or decreasing.
// A slice is monotonic if its elements from left to right are either
// non-decreasing or non-increasing.
func Monotonic(l []int) bool {
	// Slices with 0 or 1 elements are considered monotonic.
	if len(l) <= 1 {
		return true
	}

	increasing := true
	decreasing := true

	// Iterate through the slice to check both conditions in a single pass.
	for i := 1; i < len(l); i++ {
		if l[i] < l[i-1] {
			increasing = false
		}
		if l[i] > l[i-1] {
			decreasing = false
		}
	}

	// The slice is monotonic if it's either non-decreasing or non-increasing.
	return increasing || decreasing
}

// main function demonstrates the usage of Monotonic
// with the examples from the original Python code.
func main() {
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))
	fmt.Println(Monotonic([]int{1, 20, 4, 10}))
	fmt.Println(Monotonic([]int{4, 1, 0, -10}))
}
