package main

import "sort"

func Common(l1 []int, l2 []int) []int {
	// Use a map to track common elements (like Python's set)
	common := make(map[int]bool)

	// Find common elements using nested loops
	for _, e1 := range l1 {
		for _, e2 := range l2 {
			if e1 == e2 {
				common[e1] = true
			}
		}
	}

	// Convert map keys to a slice
	result := make([]int, 0, len(common))
	for k := range common {
		result = append(result, k)
	}

	// Sort the result in ascending order
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}

func main() {
	// Test examples
	println(Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121}))
	println(Common([]int{5, 3, 2, 8}, []int{3, 2}))
}
