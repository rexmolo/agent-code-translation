package main

import "fmt"

func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Check if non-decreasing (monotonic increasing)
	increasing := true
	for i := 1; i < len(l); i++ {
		if l[i] < l[i-1] {
			increasing = false
			break
		}
	}
	if increasing {
		return true
	}

	// Check if non-increasing (monotonic decreasing)
	decreasing := true
	for i := 1; i < len(l); i++ {
		if l[i] > l[i-1] {
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
	fmt.Println(Monotonic([]int{}))              // true (edge case)
	fmt.Println(Monotonic([]int{5}))             // true (edge case)
}
