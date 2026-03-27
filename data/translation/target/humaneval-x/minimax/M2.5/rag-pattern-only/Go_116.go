package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func SortArray(arr []int) []int {
	// Create a copy to avoid modifying the original slice
	result := make([]int, len(arr))
	copy(result, arr)

	sort.Slice(result, func(i, j int) bool {
		countI := bits.OnesCount(uint(result[i]))
		countJ := bits.OnesCount(uint(result[j]))

		// Primary sort: by number of 1s in binary representation
		if countI != countJ {
			return countI < countJ
		}
		// Secondary sort: by decimal value for tie-breaking
		return result[i] < result[j]
	})

	return result
}

func main() {
	// Test cases
	fmt.Println(SortArray([]int{1, 5, 2, 3, 4}))   // Expected: [1 2 3 4 5]
	fmt.Println(SortArray([]int{-2, -3, -4, -5, -6})) // Expected: [-6 -5 -4 -3 -2]
	fmt.Println(SortArray([]int{1, 0, 2, 3, 4}))   // Expected: [0 1 2 3 4]
}