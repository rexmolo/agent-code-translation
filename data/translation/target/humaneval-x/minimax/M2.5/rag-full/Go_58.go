package main

import (
	"fmt"
	"sort"
)

func Common(l1 []int, l2 []int) []int {
	// Use a map as a set to track common elements
	common := make(map[int]struct{})

	// Find common elements by comparing each element in l1 with each element in l2
	for _, e1 := range l1 {
		for _, e2 := range l2 {
			if e1 == e2 {
				common[e1] = struct{}{}
			}
		}
	}

	// Convert map keys to slice
	result := make([]int, 0, len(common))
	for k := range common {
		result = append(result, k)
	}

	// Sort the result in ascending order
	sort.Ints(result)

	return result
}

func main() {
	// Test the function
	fmt.Println(Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121}))
	fmt.Println(Common([]int{5, 3, 2, 8}, []int{3, 2}))
}
