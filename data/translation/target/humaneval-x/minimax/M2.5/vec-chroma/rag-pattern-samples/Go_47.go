package main

import (
	"fmt"
	"sort"
)

func Median(l []int) float64 {
	// Make a copy to avoid modifying the original slice
	// (matching Python's sorted() which returns a new list)
	sorted := make([]int, len(l))
	copy(sorted, l)

	sort.Ints(sorted)

	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	return (float64(sorted[n/2-1]) + float64(sorted[n/2])) / 2.0
}

func main() {
	fmt.Println(Median([]int{3, 1, 2, 4, 5}))
	fmt.Println(Median([]int{-10, 4, 6, 1000, 10, 20}))
}