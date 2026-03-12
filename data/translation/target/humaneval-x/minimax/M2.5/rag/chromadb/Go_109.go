package main

import (
	"slices"
	"sort"
)

func MoveOneBall(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	// Find minimum value and its index
	minValue := slices.Min(arr)
	minIndex := slices.Index(arr, minValue)

	// Create rotated array starting from minIndex
	rotated := append(arr[minIndex:], arr[:minIndex]...)

	// Create sorted copy to compare against
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Slice(sortedArr, func(i, j int) bool { return sortedArr[i] < sortedArr[j] })

	// Compare rotated array with sorted array
	for i := range rotated {
		if rotated[i] != sortedArr[i] {
			return false
		}
	}
	return true
}
