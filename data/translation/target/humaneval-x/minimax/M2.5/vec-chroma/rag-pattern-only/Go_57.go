package main

import "fmt"

func Monotonic(l []int) bool {
	// Empty or single-element slice is monotonic
	if len(l) <= 1 {
		return true
	}

	// Check if non-decreasing (allowing equal elements)
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

	// Check if non-increasing (allowing equal elements)
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
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))   // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10}))  // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10}))  // true
}