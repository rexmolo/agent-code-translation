package main

import (
	"fmt"
	"strconv"
)

// AddElements calculates the sum of elements with at most two digits
// from the first k elements of arr.
// It corresponds to the Python function add_elements.
func AddElements(arr []int, k int) int {
	// Given the constraint `1 <= k <= len(arr)`, we don't need to check
	// if k is out of bounds for slicing.
	// The slice `arr[:k]` will include elements from index 0 up to, but not including, k.

	sum := 0
	for _, elem := range arr[:k] {
		// In Go, we can check the number of digits by converting the integer to a string
		// using strconv.Itoa and then checking its length. This directly mimics
		// the Python logic `len(str(elem))`.
		if len(strconv.Itoa(elem)) <= 2 {
			sum += elem
		}
	}
	return sum
}

func main() {
	// Example from the problem description
	arr := []int{111, 21, 3, 4000, 5, 6, 7, 8, 9}
	k := 4
	result := AddElements(arr, k)
	fmt.Println(result) // Expected output: 24 (sum of 21 + 3)

	// Another example
	arr2 := []int{1, 10, 100, 1000}
	k2 := 4
	result2 := AddElements(arr2, k2)
	fmt.Println(result2) // Expected output: 11 (sum of 1 + 10)
}
