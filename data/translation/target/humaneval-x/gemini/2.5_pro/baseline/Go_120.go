package main

import (
	"fmt"
	"sort"
)

// Maximum given an array arr of integers and a positive integer k, returns a sorted list
// of length k with the maximum k numbers in arr.
func Maximum(arr []int, k int) []int {
	// The Python code sorts the list in-place. The Go function also modifies
	// the input slice `arr` because sort.Ints sorts in-place.
	// If the caller needs to preserve the original slice, they should pass a copy.

	if k == 0 {
		return []int{}
	}

	// sort.Ints sorts the slice in ascending order.
	sort.Ints(arr)

	// Slicing from len(arr)-k to the end gives the k largest elements.
	ans := arr[len(arr)-k:]
	return ans
}

// main function to demonstrate the Maximum function with examples.
func main() {
	// Example 1
	arr1 := []int{-3, -4, 5}
	k1 := 3
	fmt.Printf("Input: arr = %v, k = %d\n", arr1, k1)
	// Note: a new slice is created to avoid modifying the original for demonstration clarity.
	arr1Copy := make([]int, len(arr1))
	copy(arr1Copy, arr1)
	result1 := Maximum(arr1Copy, k1)
	fmt.Printf("Output: %v\n\n", result1)

	// Example 2
	arr2 := []int{4, -4, 4}
	k2 := 2
	fmt.Printf("Input: arr = %v, k = %d\n", arr2, k2)
	arr2Copy := make([]int, len(arr2))
	copy(arr2Copy, arr2)
	result2 := Maximum(arr2Copy, k2)
	fmt.Printf("Output: %v\n\n", result2)

	// Example 3
	arr3 := []int{-3, 2, 1, 2, -1, -2, 1}
	k3 := 1
	fmt.Printf("Input: arr = %v, k = %d\n", arr3, k3)
	arr3Copy := make([]int, len(arr3))
	copy(arr3Copy, arr3)
	result3 := Maximum(arr3Copy, k3)
	fmt.Printf("Output: %v\n\n", result3)
}
