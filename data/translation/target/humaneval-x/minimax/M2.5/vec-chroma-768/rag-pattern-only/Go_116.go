package main

import (
    "cmp"
    "slices"
    "math/bits"
)

// SortArray sorts an array of integers by the number of ones in their
// binary representation (ascending), then by value (ascending) for equal popcounts.
func SortArray(arr []int) []int {
    result := slices.Clone(arr)
    slices.SortFunc(result, func(a, b int) int {
        // Compare by popcount first (number of 1s in binary)
        pcA := bits.OnesCount64(uint64(a))
        pcB := bits.OnesCount64(uint64(b))
        if pcA != pcB {
            return pcA - pcB
        }
        // If popcounts are equal, sort by decimal value
        return cmp.Compare(a, b)
    })
    return result
}
