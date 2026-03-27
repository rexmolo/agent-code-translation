package main

import (
	"sort"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	// Create a copy to avoid modifying the input slice
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	return arrCopy[len(arrCopy)-k:]
}