package main

import "sort"

// MoveOneBall determines if an array can be sorted into non-decreasing order
// by performing any number of right-shift operations.
func MoveOneBall(arr []int) bool {
	// If the given array is empty, it's considered sorted.
	if len(arr) == 0 {
		return true
	}

	// Create a sorted version of the array for comparison.
	// A copy is made to avoid modifying the original input slice.
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// If an array can be sorted by rotation, it must be a rotated version
	// of the sorted array. By rotating the array to put its smallest element
	// at the front, we can check if it matches the sorted version.

	// Find the index of the minimum element in the original array.
	// Go does not have a built-in `min` or `index` function for slices like Python.
	if len(arr) == 1 {
		return true // A single element array is always sorted
	}
	minIndex := 0
	// No need to check arr[0] against itself, start loop at 1.
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	// Create a new array by rotating the original array so the minimum
	// element is at the start. This corresponds to Python's
	// `arr[min_index:] + arr[0:min_index]`.
	rotatedArr := append(arr[minIndex:], arr[:minIndex]...)

	// Compare the rotated array with the sorted version element by element.
	for i := 0; i < len(arr); i++ {
		if rotatedArr[i] != sortedArr[i] {
			return false
		}
	}

	return true
}
