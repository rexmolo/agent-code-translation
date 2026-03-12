package main

import (
	"fmt"
	"sort"
)

func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Create copies to sort without modifying original
	ascending := make([]int, len(l))
	descending := make([]int, len(l))
	copy(ascending, l)
	copy(descending, l)

	// Sort ascending
	sort.Ints(ascending)

	// Sort descending
	sort.Sort(sort.Reverse(sort.IntSlice(descending)))

	// Check if original equals ascending or descending
	for i := range l {
		if l[i] != ascending[i] && l[i] != descending[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))   // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10}))  // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10})) // true
}