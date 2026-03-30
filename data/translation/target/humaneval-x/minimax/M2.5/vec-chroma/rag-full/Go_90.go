package main

import (
	"fmt"
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Remove duplicates using a map (Go's equivalent of set)
	uniqueMap := make(map[int]struct{})
	for _, v := range lst {
		uniqueMap[v] = struct{}{}
	}

	// Convert map keys to slice for sorting
	unique := make([]int, 0, len(uniqueMap))
	for v := range uniqueMap {
		unique = append(unique, v)
	}

	// Return nil if there are fewer than 2 unique elements
	if len(unique) < 2 {
		return nil
	}

	// Sort the unique elements in ascending order
	sort.Ints(unique)

	// Return the second smallest element (index 1)
	return unique[1]
}

func main() {
	// Test cases to verify correctness
	fmt.Println(NextSmallest([]int{1, 2, 3, 4, 5})) // Expected: 2
	fmt.Println(NextSmallest([]int{5, 1, 4, 3, 2})) // Expected: 2
	fmt.Println(NextSmallest([]int{}))              // Expected: nil
	fmt.Println(NextSmallest([]int{1, 1}))          // Expected: nil
}