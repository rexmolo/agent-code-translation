package main

import (
	"fmt"
	"sort"
)

func SortArray(array []int) []int {
	// Handle empty array - return empty slice
	if len(array) == 0 {
		return []int{}
	}

	// Make a copy to avoid modifying the original array
	result := make([]int, len(array))
	copy(result, array)

	// Handle single element - already sorted
	if len(array) == 1 {
		return result
	}

	// Determine sort direction: (first + last) odd -> ascending, even -> descending
	sum := array[0] + array[len(array)-1]
	if sum%2 == 0 {
		// Even sum -> sort in descending order
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
	} else {
		// Odd sum -> sort in ascending order
		sort.Ints(result)
	}

	return result
}

func main() {
	// Test cases from the examples
	testCases := [][]int{
		{},
		{5},
		{2, 4, 3, 0, 1, 5},
		{2, 4, 3, 0, 1, 5, 6},
	}

	for _, tc := range testCases {
		result := SortArray(tc)
		fmt.Printf("sort_array(%v) => %v\n", tc, result)
	}
}
