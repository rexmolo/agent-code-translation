package main

import (
	"fmt"
	"sort"
)

func SortArray(array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	// Create a copy to avoid modifying the original array
	result := make([]int, len(array))
	copy(result, array)

	// Determine sort direction based on whether sum of first and last is even
	descending := (array[0]+array[len(array)-1])%2 == 0

	if descending {
		// Sort in descending order
		sort.Slice(result, func(i, j int) bool {
			return result[i] > result[j]
		})
	} else {
		// Sort in ascending order
		sort.Ints(result)
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(SortArray([]int{}))           // []
	fmt.Println(SortArray([]int{5}))           // [5]
	fmt.Println(SortArray([]int{2, 4, 3, 0, 1, 5})) // [0 1 2 3 4 5]
	fmt.Println(SortArray([]int{2, 4, 3, 0, 1, 5, 6})) // [6 5 4 3 2 1 0]
}
