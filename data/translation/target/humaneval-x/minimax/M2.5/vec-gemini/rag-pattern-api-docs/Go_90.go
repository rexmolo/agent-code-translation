package main

import (
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Create a set (map with empty struct value) to remove duplicates
	set := make(map[int]struct{})
	for _, v := range lst {
		set[v] = struct{}{}
	}

	// Convert map keys to slice
	unique := make([]int, 0, len(set))
	for k := range set {
		unique = append(unique, k)
	}

	// Sort the slice in ascending order
	sort.Slice(unique, func(i, j int) bool {
		return unique[i] < unique[j]
	})

	// Return the 2nd smallest element or nil if not enough elements
	if len(unique) < 2 {
		return nil
	}
	return unique[1]
}

func main() {
	// Test cases
	_ = NextSmallest([]int{1, 2, 3, 4, 5})
	_ = NextSmallest([]int{5, 1, 4, 3, 2})
	_ = NextSmallest([]int{})
	_ = NextSmallest([]int{1, 1})
}
