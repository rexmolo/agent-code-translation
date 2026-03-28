package main

import (
	"fmt"
	"sort"
)

func SortArray(arr []int) []int {
	// First sort by value (ascending)
	sort.Ints(arr)

	// Then sort by number of 1's in binary representation (ascending)
	// Using stable sort to maintain value order for equal bit counts
	sort.SliceStable(arr, func(i, j int) bool {
		countI := bitCount(arr[i])
		countJ := bitCount(arr[j])
		if countI == countJ {
			return arr[i] < arr[j]
		}
		return countI < countJ
	})

	return arr
}

// bitCount returns the number of 1's in the binary representation of n
func bitCount(n int) int {
	if n < 0 {
		// For negative numbers, use two's complement representation
		// Convert to unsigned to get the full bit pattern
		n = ^n + 1
	}
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// For testing
func main() {
	// Test cases
	fmt.Println(SortArray([]int{1, 5, 2, 3, 4}))   // Expected: [1, 2, 3, 4, 5] or [1, 2, 4, 3, 5]?
	fmt.Println(SortArray([]int{-2, -3, -4, -5, -6})) // Expected: [-6, -5, -4, -3, -2]
	fmt.Println(SortArray([]int{1, 0, 2, 3, 4}))    // Expected: [0, 1, 2, 3, 4]
}
