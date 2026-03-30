package main

import (
	"fmt"
	"sort"
)

func SortArray(array []int) []int {
	// Handle empty array case - return empty slice
	if len(array) == 0 {
		return []int{}
	}

	// Create a copy to avoid modifying the original array
	result := make([]int, len(array))
	copy(result, array)

	// Determine sort direction based on sum of first and last elements
	// Even sum -> descending, Odd sum -> ascending
	sum := array[0] + array[len(array)-1]
	if sum%2 == 0 {
		// Even sum: sort in descending order
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
	} else {
		// Odd sum: sort in ascending order
		sort.Slice(result, func(i, j int) bool {
			return result[i] < result[j]
		})
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(SortArray([]int{}))           // []
	fmt.Println(SortArray([]int{5}))          // [5]
	fmt.Println(SortArray([]int{2, 4, 3, 0, 1, 5})) // [0 1 2 3 4 5]
	fmt.Println(SortArray([]int{2, 4, 3, 0, 1, 5, 6})) // [6 5 4 3 2 1 0]
}
