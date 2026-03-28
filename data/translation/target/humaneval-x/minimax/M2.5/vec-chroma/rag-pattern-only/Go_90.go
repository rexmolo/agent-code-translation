package main

import (
	"fmt"
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Remove duplicates by using a map to track unique values
	unique := make(map[int]bool)
	for _, v := range lst {
		unique[v] = true
	}
	
	// Convert map keys to a slice
	sorted := make([]int, 0, len(unique))
	for k := range unique {
		sorted = append(sorted, k)
	}
	
	// Sort the slice in ascending order
	sort.Ints(sorted)
	
	// Return the 2nd smallest element or nil if there aren't enough unique elements
	if len(sorted) < 2 {
		return nil
	}
	return sorted[1]
}

func main() {
	// Test cases
	fmt.Println(NextSmallest([]int{1, 2, 3, 4, 5})) // 2
	fmt.Println(NextSmallest([]int{5, 1, 4, 3, 2})) // 2
	fmt.Println(NextSmallest([]int{}))              // <nil>
	fmt.Println(NextSmallest([]int{1, 1}))          // <nil>
}
