package main

import (
	"sort"
)

func Common(l1 []int, l2 []int) []int {
	ret := make(map[int]bool)
	for _, e1 := range l1 {
		for _, e2 := range l2 {
			if e1 == e2 {
				ret[e1] = true
			}
		}
	}

	// Convert map keys to slice
	result := make([]int, 0, len(ret))
	for k := range ret {
		result = append(result, k)
	}

	// Sort the result
	sort.Ints(result)
	return result
}
