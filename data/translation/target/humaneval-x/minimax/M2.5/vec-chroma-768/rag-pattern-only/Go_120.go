package main

import (
	"fmt"
	"sort"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	sort.Ints(arr)
	return arr[len(arr)-k:]
}

func main() {
	// Test cases
	testCases := []struct {
		arr []int
		k   int
	}{
		{[]int{-3, -4, 5}, 3},
		{[]int{4, -4, 4}, 2},
		{[]int{-3, 2, 1, 2, -1, -2, 1}, 1},
	}

	for _, tc := range testCases {
		// Create a copy to preserve original for display
		original := make([]int, len(tc.arr))
		copy(original, tc.arr)
		result := Maximum(tc.arr, tc.k)
		fmt.Printf("arr = %v, k = %d, result = %v\n", original, tc.k, result)
	}
}
