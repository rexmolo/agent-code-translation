package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortArray(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		countI := strings.Count(fmt.Sprintf("%b", abs(arr[i])), "1")
		countJ := strings.Count(fmt.Sprintf("%b", abs(arr[j])), "1")
		if countI == countJ {
			return arr[i] < arr[j]
		}
		return countI < countJ
	})
	return arr
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	// Test cases
	fmt.Println(SortArray([]int{1, 5, 2, 3, 4}))   // Expected: [1, 2, 4, 5, 3]
	fmt.Println(SortArray([]int{-2, -3, -4, -5, -6})) // Expected: [-6, -5, -4, -3, -2]
	fmt.Println(SortArray([]int{1, 0, 2, 3, 4}))    // Expected: [0, 1, 2, 4, 3]
}