package main

import (
	"fmt"
	"sort"
)

func Monotonic(l []int) bool {
	if len(l) <= 1 {
		return true
	}

	// Make copies to avoid modifying original slice
	asc := make([]int, len(l))
	copy(asc, l)

	desc := make([]int, len(l))
	copy(desc, l)

	// Sort ascending and compare
	sort.Ints(asc)
	if equal(l, asc) {
		return true
	}

	// Sort descending and compare
	sort.Sort(sort.Reverse(sort.IntSlice(desc)))
	return equal(l, desc)
}

// equal checks if two slices have the same elements
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(Monotonic([]int{1, 2, 4, 20}))   // true
	fmt.Println(Monotonic([]int{1, 20, 4, 10}))   // false
	fmt.Println(Monotonic([]int{4, 1, 0, -10}))   // true
}
