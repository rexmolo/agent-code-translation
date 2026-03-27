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

	// Find minimum value in the array
	minValue := arr[0]
	for _, v := range arr {
		if v < minValue {
			minValue = v
		}
	}

	// Find index of minimum value
	minIndex := slices.Index(arr, minValue)

	// Create rotated array: elements from minIndex to end, then from start to minIndex
	myArr := append(arr[minIndex:], arr[0:minIndex]...)

	// Create sorted copy of the array
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Compare element by element
	for i := 0; i < len(arr); i++ {
		if myArr[i] != sortedArr[i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(MoveOneBall([]int{3, 4, 5, 1, 2})) // true
	fmt.Println(MoveOneBall([]int{3, 5, 4, 1, 2})) // false
	fmt.Println(MoveOneBall([]int{}))              // true
	fmt.Println(MoveOneBall([]int{1}))              // true
	fmt.Println(MoveOneBall([]int{2, 1}))           // true
}
