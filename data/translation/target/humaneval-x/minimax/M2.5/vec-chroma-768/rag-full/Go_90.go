package main

import (
	"fmt"
	"sort"
)

func NextSmallest(lst []int) interface{} {
	// Remove duplicates using a map
	seen := make(map[int]bool)
	unique := make([]int, 0)
	for _, v := range lst {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	
	// Sort the unique elements
	sort.Ints(unique)
	
	// Return the 2nd smallest element or nil if not enough elements
	if len(unique) < 2 {
		return nil
	}
	return unique[1]
}

func main() {
	// Test cases
	fmt.Println(NextSmallest([]int{1, 2, 3, 4, 5})) // 2
	fmt.Println(NextSmallest([]int{5, 1, 4, 3, 2})) // 2
	fmt.Println(NextSmallest([]int{}))             // <nil>
	fmt.Println(NextSmallest([]int{1, 1}))         // <nil>
}