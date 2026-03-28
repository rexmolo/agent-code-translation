package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortArray(arr []int) []int {
	// Make a copy to avoid modifying the original slice
	result := make([]int, len(arr))
	copy(result, arr)

	// Sort by count of 1s in binary representation (ascending),
	// then by decimal value for equal counts
	sort.Slice(result, func(i, j int) bool {
		countI := countOnes(result[i])
		countJ := countOnes(result[j])
		if countI == countJ {
			return result[i] < result[j]
		}
		return countI < countJ
	})

	return result
}

// countOnes returns the number of 1s in the binary representation of n
func countOnes(n int) int {
	// Convert to binary string and count '1' characters
	binStr := fmt.Sprintf("%b", n)
	return strings.Count(binStr, "1")
}

func main() {
	// Test cases from docstring
	fmt.Println(SortArray([]int{1, 5, 2, 3, 4}))
	fmt.Println(SortArray([]int{-2, -3, -4, -5, -6}))
	fmt.Println(SortArray([]int{1, 0, 2, 3, 4}))
}
