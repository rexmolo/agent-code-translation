package main

import (
	"sort"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	// Make a copy to avoid modifying the original slice
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	sort.Ints(copyArr)
	return copyArr[len(copyArr)-k:]
}
