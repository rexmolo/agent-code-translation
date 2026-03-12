package main

import (
	"sort"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	// Create a copy to avoid modifying the original slice
	sorted := make([]int, len(arr))
	copy(sorted, arr)

	// Sort the copy in ascending order
	sort.Ints(sorted)

	// Return the last k elements (the k largest values, sorted ascending)
	return sorted[len(sorted)-k:]
}

// For testing
func main() {
	// Example tests
	arr1 := []int{-3, -4, 5}
	k1 := 3
	println(Maximum(arr1, k1)) // Expected: [-4, -3, 5]

	arr2 := []int{4, -4, 4}
	k2 := 2
	println(Maximum(arr2, k2)) // Expected: [4, 4]

	arr3 := []int{-3, 2, 1, 2, -1, -2, 1}
	k3 := 1
	println(Maximum(arr3, k3)) // Expected: [2]
}
