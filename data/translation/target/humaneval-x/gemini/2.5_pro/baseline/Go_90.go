package main

import (
	"sort"
)

// NextSmallest finds the 2nd smallest unique element in a slice of integers.
// It returns nil if there is no such element.
func NextSmallest(lst []int) interface{} {
	// Use a map to find unique elements, mimicking Python's set().
	uniqueKeys := make(map[int]bool)
	for _, item := range lst {
		uniqueKeys[item] = true
	}

	// If there are fewer than 2 unique elements, there can't be a 2nd smallest.
	if len(uniqueKeys) < 2 {
		return nil
	}

	// Create a new slice from the unique keys.
	uniqueList := make([]int, 0, len(uniqueKeys))
	for key := range uniqueKeys {
		uniqueList = append(uniqueList, key)
	}

	// Sort the slice of unique elements.
	sort.Ints(uniqueList)

	// Return the 2nd smallest element.
	return uniqueList[1]
}
