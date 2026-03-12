package main

import (
	"math"
)

// Pluck finds the node with the smallest even value in a slice of non-negative integers.
// It returns a slice containing the value and its index, e.g., [value, index].
// If multiple nodes have the same smallest even value, it chooses the one with the smallest index.
// If the input slice is empty or contains no even numbers, it returns an empty slice.
func Pluck(arr []int) []int {
	// Initialize with a value that indicates no even number has been found yet.
	resultIndex := -1
	// Start with the largest possible integer to find the minimum.
	smallestEven := math.MaxInt

	// Iterate through the slice once to find the smallest even number and its index.
	for i, val := range arr {
		// Check if the current value is even.
		if val%2 == 0 {
			// If this even value is smaller than the smallest one found so far,
			// update the smallest value and its index.
			// The strict '<' ensures that for ties, the one with the first (smallest) index is kept.
			if val < smallestEven {
				smallestEven = val
				resultIndex = i
			}
		}
	}

	// If resultIndex is still -1, no even numbers were found.
	if resultIndex == -1 {
		return []int{}
	}

	// Return the smallest even value and its original index.
	return []int{smallestEven, resultIndex}
}
