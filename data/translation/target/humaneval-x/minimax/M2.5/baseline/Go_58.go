package main

import (
	"fmt"
	"sort"
)

func Common(l1 []int, l2 []int) []int {
	// Use a map to track unique common elements
	commonMap := make(map[int]bool)

	// Find common elements using nested loops
	for _, e1 := range l1 {
		for _, e2 := range l2 {
			if e1 == e2 {
				commonMap[e1] = true
			}
		}
	}

	// Convert map keys to slice
	ret := make([]int, 0, len(commonMap))
	for k := range commonMap {
		ret = append(ret, k)
	}

	// Sort the slice in ascending order
	sort.Ints(ret)

	return ret
}

func main() {
	// Test cases from docstring
	fmt.Println(Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121}))
	fmt.Println(Common([]int{5, 3, 2, 8}, []int{3, 2}))
}
