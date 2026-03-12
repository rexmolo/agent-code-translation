package main

import (
	"sort"
)

// SortEven takes a slice l and returns a new slice l' such that
// l' is identical to l in the odd indicies, while its values at the even indicies are equal
// to the values of the even indicies of l, but sorted.
func SortEven(l []int) []int {
	// Step 1: Extract elements at even indices into a new slice.
	// The number of even elements is len(l)/2 if len(l) is even, and len(l)/2 + 1 if it's odd.
	// (len(l) + 1) / 2 correctly calculates this for both cases.
	eventsCapacity := (len(l) + 1) / 2
	evens := make([]int, 0, eventsCapacity)
	for i := 0; i < len(l); i += 2 {
		evens = append(evens, l[i])
	}

	// Step 2: Sort the slice of even-indexed elements.
	sort.Ints(evens)

	// Step 3: Create a result slice as a copy of the original.
	// This preserves the original odd-indexed elements.
	result := make([]int, len(l))
	copy(result, l)

	// Step 4: Overwrite the even indices in the result slice with the sorted elements.
	for i, val := range evens {
		result[i*2] = val
	}

	return result
}
