package main

import (
	"fmt"
	"sort"
)

func MoveOneBall(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	// Find the minimum value and its index
	minValue := arr[0]
	minIndex := 0
	for i, v := range arr {
		if v < minValue {
			minValue = v
			minIndex = i
		}
	}

	// Create sorted copy to compare against
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Create rotated array by taking elements from minIndex to end,
	// then appending elements from start to minIndex
	rotated := append(arr[minIndex:], arr[0:minIndex]...)

	// Compare rotated array with sorted array
	for i := range rotated {
		if rotated[i] != sortedArr[i] {
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
	fmt.Println(MoveOneBall([]int{5, 1, 2, 3, 4})) // true
}