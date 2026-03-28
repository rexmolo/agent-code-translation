package main

import (
	"fmt"
	"sort"
)

func MoveOneBall(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	// Create a sorted copy
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

	// Create rotated array starting from minIndex
	myArr := make([]int, len(arr))
	copy(myArr, arr[minIndex:])
	copy(myArr[len(arr)-minIndex:], arr[:minIndex])

	// Compare rotated array with sorted array
	for i := 0; i < len(arr); i++ {
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
	fmt.Println(MoveOneBall([]int{}))               // Expected: true
}
