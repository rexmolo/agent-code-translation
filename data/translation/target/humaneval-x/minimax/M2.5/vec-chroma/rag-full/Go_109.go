package main

import (
	"fmt"
	"sort"
)

func MoveOneBall(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	// Find minimum value and its index
	minValue := arr[0]
	minIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < minValue {
			minValue = arr[i]
			minIndex = i
		}
	}

	// Rotate array so that min element comes first (simulating right shifts)
	myArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		myArr[i] = arr[(minIndex+i)%len(arr)]
	}

	// Create sorted copy for comparison
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Compare rotated array with sorted array
	for i := 0; i < len(arr); i++ {
		if myArr[i] != sortedArr[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test examples
	fmt.Println(MoveOneBall([]int{3, 4, 5, 1, 2})) // true
	fmt.Println(MoveOneBall([]int{3, 5, 4, 1, 2})) // false
	fmt.Println(MoveOneBall([]int{}))              // true
	fmt.Println(MoveOneBall([]int{1}))             // true
}