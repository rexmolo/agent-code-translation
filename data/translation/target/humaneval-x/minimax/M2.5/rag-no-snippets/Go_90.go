package main

import (
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Remove duplicates using a map to simulate set(lst)
	seen := make(map[int]bool)
	unique := []int{}
	for _, v := range lst {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}

	// Sort the unique elements (equivalent to sorted(set(lst)))
	sort.Ints(unique)

	// Return the 2nd smallest element or nil if there are fewer than 2
	if len(unique) < 2 {
		return nil
	}
	return unique[1]
}