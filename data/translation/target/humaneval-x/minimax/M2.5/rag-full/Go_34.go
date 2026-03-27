package main

import "slices"

func Unique(l []int) []int {
	// Use map to simulate a set for unique elements
	uniqueMap := make(map[int]struct{})
	for _, v := range l {
		uniqueMap[v] = struct{}{}
	}

	// Convert map keys to slice
	result := make([]int, 0, len(uniqueMap))
	for k := range uniqueMap {
		result = append(result, k)
	}

	// Sort the slice using slices.Sort (equivalent to Python's sorted())
	slices.Sort(result)

	return result
}