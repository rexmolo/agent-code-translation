package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortArray(arr []int) []int {
    if arr == nil {
        return nil
    }
    // First sort by decimal value
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    // Then sort by count of 1s in binary representation (stable to preserve value order for ties)
    sort.SliceStable(arr, func(i, j int) bool {
        return countOnes(arr[i]) < countOnes(arr[j])
    })
    return arr
}

func countOnes(n int) int {
    // Use fmt.Sprintf to get binary representation, skip the "0b" prefix
    binary := fmt.Sprintf("%b", n)
    return strings.Count(binary, "1")
}
