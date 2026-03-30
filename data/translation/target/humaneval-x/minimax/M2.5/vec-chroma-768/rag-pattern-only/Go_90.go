package main

import (
	"fmt"
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Create a map to track unique elements (removes duplicates)
	unique := make(map[int]bool)
	for _, n := range lst {
		unique[n] = true
	}

	// Convert map keys to a slice
	uniqueVals := make([]int, 0, len(unique))
	for k := range unique {
		uniqueVals = append(uniqueVals, k)
	}

	// Sort the slice in ascending order
	sort.Ints(uniqueVals)

	// Return the 2nd smallest element or nil if fewer than 2 unique elements
	if len(uniqueVals) < 2 {
		return nil
	}
	return uniqueVals[1]
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{5, 1, 4, 3, 2},
		{},
		{1, 1},
	}

	for _, lst := range testCases {
		result := NextSmallest(lst)
		fmt.Printf("NextSmallest(%v) = %v\n", lst, result)
	}
}
