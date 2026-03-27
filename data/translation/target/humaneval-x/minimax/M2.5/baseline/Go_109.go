package main

import (
	"sort"
)

func MoveOneBall(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	// Create a sorted copy to compare against
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Find minimum value and its index
	minValue := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	// Create rotated array starting from minIndex (simulating right shifts)
	myArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		myArr[i] = arr[(minIndex+i)%len(arr)]
	}

	// Compare rotated array with sorted array
	for i := 0; i < len(arr); i++ {
		if myArr[i] != sortedArr[i] {
			return false
		}
	}
	return true
}
