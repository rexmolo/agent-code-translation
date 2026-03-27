package main

import (
	"fmt"
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Create a set to remove duplicates
	unique := make(map[int]struct{})
	for _, v := range lst {
		unique[v] = struct{}{}
	}

	// Convert set to slice
	s := make([]int, 0, len(unique))
	for v := range unique {
		s = append(s, v)
	}

	// Sort the slice
	sort.Ints(s)

	// Return the second smallest or nil if there are fewer than 2 elements
	if len(s) < 2 {
		return nil
	}
	return s[1]
}

func main() {
	// Test cases
	fmt.Println(NextSmallest([]int{1, 2, 3, 4, 5})) // 2
	fmt.Println(NextSmallest([]int{5, 1, 4, 3, 2})) // 2
	fmt.Println(NextSmallest([]int{}))               // nil
	fmt.Println(NextSmallest([]int{1, 1}))           // nil
}
