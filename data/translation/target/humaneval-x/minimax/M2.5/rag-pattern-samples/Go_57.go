package main

import "fmt"

// Monotonic returns true if the slice elements are monotonically
// increasing or decreasing.
func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Check if monotonically increasing
	increasing := true
	for i := 0; i < len(l)-1; i++ {
		if l[i] > l[i+1] {
			increasing = false
			break
		}
	}
	if increasing {
		return true
	}

	// Check if monotonically decreasing
	decreasing := true
	for i := 0; i < len(l)-1; i++ {
		if l[i] < l[i+1] {
			decreasing = false
			break
		}
	}
	return decreasing
}

func main() {
	// Test cases
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))  // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10})) // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10})) // true
}
