package main

import (
	"sort"
)

func Median(l []int) float64 {
	sort.Ints(l)
	n := len(l)
	mid := n / 2
	if n%2 == 1 {
		return float64(l[mid])
	}
	return float64(l[mid-1]+l[mid]) / 2.0
}
