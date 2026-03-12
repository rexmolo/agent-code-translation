package main

import "fmt"

// CanArrange returns the largest index of an element which
// is not greater than or equal to the element immediately preceding it. If
// no such element exists then it returns -1. The given array will not contain
// duplicate values.
//
// Examples:
// CanArrange([]int{1, 2, 4, 3, 5}) = 3
// CanArrange([]int{1, 2, 3}) = -1
func CanArrange(arr []int) int {
	ind := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			ind = i
		}
	}
	return ind
}

func main() {
	// Example 1
	arr1 := []int{1, 2, 4, 3, 5}
	fmt.Printf("Input: %v, Output: %d\n", arr1, CanArrange(arr1))

	// Example 2
	arr2 := []int{1, 2, 3}
	fmt.Printf("Input: %v, Output: %d\n", arr2, CanArrange(arr2))
}
