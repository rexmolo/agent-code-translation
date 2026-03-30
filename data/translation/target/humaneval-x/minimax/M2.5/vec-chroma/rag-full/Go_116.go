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

	// First, sort by value (ascending) - this is the secondary sort key
	sort.Ints(result)

	// Then, stable sort by number of 1-bits in binary (ascending)
	// This preserves the relative order of elements with equal bit counts
	sort.SliceStable(result, func(i, j int) bool {
		return onesCount(result[i]) < onesCount(result[j])
	})

	return result
}

// onesCount returns the number of 1s in the binary representation
func onesCount(n int) int {
	if n >= 0 {
		return bits.OnesCount(uint(n))
	}
	// For negative numbers, use two's complement representation
	return bits.OnesCount(^uint(n))
}

func main() {
	// Test cases from the docstring
	testCases := [][]int{
		{1, 5, 2, 3, 4},
		{-2, -3, -4, -5, -6},
		{1, 0, 2, 3, 4},
	}

	for _, tc := range testCases {
		fmt.Printf("Input:  %v\n", tc)
		fmt.Printf("Output: %v\n\n", SortArray(tc))
	}
}
