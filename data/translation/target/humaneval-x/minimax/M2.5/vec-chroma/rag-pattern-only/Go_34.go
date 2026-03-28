package main

import (
    "sort"
)

func Unique(l []int) []int {
    // Use a map to track unique elements
    seen := make(map[int]bool)
    for _, v := range l {
        seen[v] = true
    }

    // Collect unique values into a slice
    result := make([]int, 0, len(seen))
    for k := range seen {
        result = append(result, k)
    }

    // Sort and return
    sort.Ints(result)
    return result
}