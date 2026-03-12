package main

import (
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Create a map to get unique elements (equivalent to set())
	unique := make(map[int]bool)
	for _, v := range lst {
		unique[v] = true
	}

	// Convert map keys to a slice
	s := make([]int, 0, len(unique))
	for k := range unique {
		s = append(s, k)
	}

	// Sort the slice (equivalent to sorted())
	sort.Ints(s)

	// Return the 2nd smallest element or nil if less than 2 elements
	if len(s) < 2 {
		return nil
	}
	return s[1]
}

// For testing purposes
func main() {
	// Test cases
	_ = NextSmallest([]int{1, 2, 3, 4, 5})
	_ = NextSmallest([]int{5, 1, 4, 3, 2})
	_ = NextSmallest([]int{})
	_ = NextSmallest([]int{1, 1})
}
