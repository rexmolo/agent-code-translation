package main

import "slices"

func Pluck(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// Filter even numbers
	var evens []int
	for _, v := range arr {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}

	// Return empty if no even values
	if len(evens) == 0 {
		return []int{}
	}

	// Find minimum even value and its index in original array
	minEven := slices.Min(evens)
	index := slices.Index(arr, minEven)

	return []int{minEven, index}
}