package main

import (
	"math/bits"
	"sort"
)

// SortArray sorts a slice of integers based on the number of set bits in their binary representation,
// with a secondary sort based on their decimal value.
func SortArray(arr []int) []int {
	// Create a copy to avoid modifying the original slice, mimicking Python's sorted().
	result := make([]int, len(arr))
	copy(result, arr)

	// Use sort.Slice with a custom comparison function that implements the two-level sorting logic.
	sort.Slice(result, func(i, j int) bool {
		a := result[i]
		b := result[j]

		// The Python `bin(x)[2:].count('1')` logic on negative numbers
		// effectively operates on the absolute value.
		var aAbs, bAbs uint
		if a < 0 {
			aAbs = uint(-a)
		} else {
			aAbs = uint(a)
		}
		if b < 0 {
			bAbs = uint(-b)
		} else {
			bAbs = uint(b)
		}

		countA := bits.OnesCount(aAbs)
		countB := bits.OnesCount(bAbs)

		// Primary sort criterion: number of ones (ascending).
		if countA != countB {
			return countA < countB
		}

		// Secondary sort criterion: decimal value (ascending).
		return a < b
	})

	return result
}