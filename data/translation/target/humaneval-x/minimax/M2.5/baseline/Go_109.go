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
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	// Create rotated array: elements from minIndex to end, then from 0 to minIndex
	myArr := append(arr[minIndex:], arr[0:minIndex]...)

	// Create a sorted copy of the array
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Compare the rotated array with the sorted array
	for i := range myArr {
		if myArr[i] != sortedArr[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(MoveOneBall([]int{}))           // true
	fmt.Println(MoveOneBall([]int{3, 4, 5, 1, 2})) // true
	fmt.Println(MoveOneBall([]int{3, 5, 4, 1, 2})) // false
	fmt.Println(MoveOneBall([]int{1, 2, 3, 4, 5})) // true
}