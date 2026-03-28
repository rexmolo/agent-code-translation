package main

import (
	"cmp"
	"slices"
)

func SortArray(arr []int) []int {
	// First sort by value (secondary sort)
	slices.Sort(arr)

	// Then sort by number of 1s in binary representation (primary sort)
	slices.SortFunc(arr, func(a, b int) int {
		countA := countOnes(a)
		countB := countOnes(b)
		// Compare by number of 1s first
		if countA != countB {
			return cmp.Compare(countA, countB)
		}
		// If same count, compare by decimal value
		return cmp.Compare(a, b)
	})

	return arr
}

func countOnes(n int) int {
	// Count the number of 1s in the binary representation
	// Using bit counting: n & (n-1) clears the lowest set bit
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}
