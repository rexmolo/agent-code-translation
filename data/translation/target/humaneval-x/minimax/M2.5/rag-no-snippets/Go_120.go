package main

import (
	"sort"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	// Make a copy to avoid modifying the original input
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)

	// Sort in ascending order
	sort.Slice(copyArr, func(i, j int) bool {
		return copyArr[i] < copyArr[j]
	})

	// Return the last k elements (the k largest values)
	return copyArr[len(copyArr)-k:]
}
