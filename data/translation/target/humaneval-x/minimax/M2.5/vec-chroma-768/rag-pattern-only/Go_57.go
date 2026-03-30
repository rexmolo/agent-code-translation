package main

import (
	"fmt"
	"sort"
)

func Monotonic(l []int) bool {
	// Empty or single-element slice is monotonic
	if len(l) <= 1 {
		return true
	}

	// Check if sorted in ascending order
	ascending := sort.IntsAreSorted(l)

	// Check if sorted in descending order
	descending := true
	for i := 0; i < len(l)-1; i++ {
		if l[i] <= l[i+1] {
			descending = false
			break
		}
	}

	return ascending || descending
}

func main() {
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))   // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10}))  // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10}))  // true
}