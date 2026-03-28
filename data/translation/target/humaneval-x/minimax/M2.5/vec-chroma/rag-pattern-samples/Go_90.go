package main

import (
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Remove duplicates using a map
	uniqueMap := make(map[int]bool)
	for _, v := range lst {
		uniqueMap[v] = true
	}

	// Convert map keys to slice
	unique := make([]int, 0, len(uniqueMap))
	for k := range uniqueMap {
		unique = append(unique, k)
	}

	// Sort the slice
	sort.Ints(unique)

	// Return the 2nd smallest or nil if there are fewer than 2 unique elements
	if len(unique) < 2 {
		return nil
	}
	return unique[1]
}
