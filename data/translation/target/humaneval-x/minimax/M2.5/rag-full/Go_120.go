package main

import (
	"sort"
)

func Maximum(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	sort.Ints(arr)
	return arr[len(arr)-k:]
}