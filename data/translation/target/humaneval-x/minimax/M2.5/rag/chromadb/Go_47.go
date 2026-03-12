package main

import "sort"

func Median(l []int) float64 {
	sort.Ints(l)
	n := len(l)
	if n%2 == 1 {
		return float64(l[n/2])
	}
	return float64(l[n/2-1]+l[n/2]) / 2.0
}