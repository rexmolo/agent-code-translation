package main

import (
	"slices"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	// Sort the slice in ascending order
	slices.Sort(arr)
	// Return the last k elements
	return arr[len(arr)-k:]
}