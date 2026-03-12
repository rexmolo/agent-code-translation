package main

import (
	"fmt"
	"sort"
)

func MoveOneBall(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	// Create a sorted copy of the array
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Find minimum value and its index in the original array
	minValue := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	// Create rotated array starting from the minimum element's position
	myArr := append(arr[minIndex:], arr[:minIndex]...)

	// Compare the rotated array with the sorted array
	for i := range arr {
		if myArr[i] != sortedArr[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(MoveOneBall([]int{3, 4, 5, 1, 2})) // Expected: true
	fmt.Println(MoveOneBall([]int{3, 5, 4, 1, 2})) // Expected: false
	fmt.Println(MoveOneBall([]int{}))             // Expected: true
}
