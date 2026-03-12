package main

import (
	"fmt"
	"sort"
)

// Common returns sorted unique common elements for two integer slices.
func Common(l1 []int, l2 []int) []int {
	// Use a map to simulate a set for storing unique elements.
	// The value type is struct{} because it occupies 0 bytes.
	retSet := make(map[int]struct{})

	// Iterate through each element of the first slice.
	for _, e1 := range l1 {
		// For each element, iterate through the second slice to find matches.
		for _, e2 := range l2 {
			if e1 == e2 {
				// If a match is found, add it to our set.
				// If the element is already in the map, this does nothing.
				retSet[e1] = struct{}{}
			}
		}
	}

	// The keys of the map now hold the unique common elements.
	// We need to convert them to a slice.
	ret := make([]int, 0, len(retSet))
	for k := range retSet {
		ret = append(ret, k)
	}

	// Sort the resulting slice in ascending order.
	sort.Ints(ret)

	// Return the sorted slice of unique common elements.
	return ret
}

func main() {
	// Example 1 from Python docstring
	result1 := Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121})
	fmt.Println(result1)

	// Example 2 from Python docstring
	result2 := Common([]int{5, 3, 2, 8}, []int{3, 2})
	fmt.Println(result2)
}
