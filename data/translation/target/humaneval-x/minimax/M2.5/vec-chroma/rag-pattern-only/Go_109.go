package main

import (
	"sort"
)

func MoveOneBall(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	// Find minimum value and its index
	minValue := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	// Create rotated array: arr[min_index:] + arr[0:min_index]
	rotated := make([]int, len(arr))
	copy(rotated, arr[minIndex:])
	copy(rotated[len(arr)-minIndex:], arr[0:minIndex])

	// Create sorted copy
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Compare rotated array with sorted array
	for i := 0; i < len(arr); i++ {
		if rotated[i] != sortedArr[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	// fmt.Println(MoveOneBall([]int{3, 4, 5, 1, 2})) // Expected: true
	// fmt.Println(MoveOneBall([]int{3, 5, 4, 1, 2})) // Expected: false
	// fmt.Println(MoveOneBall([]int{}))              // Expected: true
}