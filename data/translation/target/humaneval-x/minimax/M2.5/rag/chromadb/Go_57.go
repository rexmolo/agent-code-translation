package main

import "fmt"

func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Check if sorted in ascending or descending order
	increasing := true
	decreasing := true

	for i := 1; i < len(l); i++ {
		if l[i] > l[i-1] {
			decreasing = false
		}
		if l[i] < l[i-1] {
			increasing = false
		}
	}

	return increasing || decreasing
}

func main() {
	// Test cases
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))   // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10})) // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10})) // true
}
