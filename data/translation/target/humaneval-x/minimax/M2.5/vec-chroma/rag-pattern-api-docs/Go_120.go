package main

import (
    "slices"
)

func Maximum(arr []int, k int) []int {
    if k == 0 {
        return []int{}
    }

    // Sort the array in ascending order
    slices.Sort(arr)

    // Return the last k elements (the largest k elements)
    start := len(arr) - k
    return arr[start:]
}
