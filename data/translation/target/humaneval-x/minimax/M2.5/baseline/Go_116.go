package main

import (
	"fmt"
	"sort"
)

func SortArray(arr []int) []int {
	// Sort by number of 1s in binary representation (ascending)
	// For same number of 1s, sort by decimal value (ascending)
	sort.Slice(arr, func(i, j int) bool {
		iOnes := countOnes(arr[i])
		jOnes := countOnes(arr[j])
		if iOnes == jOnes {
			return arr[i] < arr[j]
		}
		return iOnes < jOnes
	})
	return arr
}

func countOnes(n int) int {
	count := 0
	if n < 0 {
		n = -n // For negative numbers, use absolute value for bit count
	}
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(SortArray([]int{1, 5, 2, 3, 4}))   // Expected: [1 2 3 4 5]
	fmt.Println(SortArray([]int{-2, -3, -4, -5, -6})) // Expected: [-6 -5 -4 -3 -2]
	fmt.Println(SortArray([]int{1, 0, 2, 3, 4}))   // Expected: [0 1 2 3 4]
}
