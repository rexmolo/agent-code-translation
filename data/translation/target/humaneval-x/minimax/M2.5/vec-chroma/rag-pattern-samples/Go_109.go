package main

import (
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

	// Find minimum value and its index
	minValue := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	// Create rotated array: elements from minIndex to end, then from 0 to minIndex
	myArr := make([]int, len(arr))
	copy(myArr, arr[minIndex:])
	copy(myArr[len(arr)-minIndex:], arr[:minIndex])

	// Compare rotated array with sorted array
	for i := range arr {
		if myArr[i] != sortedArray[i] {
			return false
		}
	}
	return true
}