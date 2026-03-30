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

	// Find the minimum value and its index
	minValue := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	// Create rotated array (left rotation from minIndex)
	// This is equivalent to right rotation by (len(arr) - minIndex)
	rotated := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		rotated[i] = arr[(minIndex+i)%len(arr)]
	}

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
	fmt.Println(MoveOneBall([]int{3, 4, 5, 1, 2})) // Expected: true
	fmt.Println(MoveOneBall([]int{3, 5, 4, 1, 2})) // Expected: false
	fmt.Println(MoveOneBall([]int{}))              // Expected: true
	fmt.Println(MoveOneBall([]int{1, 2, 3, 4, 5})) // Expected: true
	fmt.Println(MoveOneBall([]int{5, 4, 3, 2, 1})) // Expected: false
}
