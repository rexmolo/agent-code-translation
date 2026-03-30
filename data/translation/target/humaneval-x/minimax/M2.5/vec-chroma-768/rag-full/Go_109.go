package main

import (
	"fmt"
	"slices"
	"sort"
)

func MoveOneBall(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	// Create a sorted copy of the array
	sortedArray := make([]int, len(arr))
	copy(sortedArray, arr)
	sort.Ints(sortedArray)

	// Find minimum value and its index in the original array
	minValue := slices.Min(arr)
	minIndex := slices.Index(arr, minValue)

	// Create rotated array starting from minIndex (this simulates right shifts)
	myArr := append(arr[minIndex:], arr[:minIndex]...)

	// Compare rotated array with sorted array
	for i := range arr {
		if myArr[i] != sortedArray[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(MoveOneBall([]int{3, 4, 5, 1, 2})) // Expected: true
	fmt.Println(MoveOneBall([]int{3, 5, 4, 1, 2})) // Expected: false
	fmt.Println(MoveOneBall([]int{}))              // Expected: true
}
