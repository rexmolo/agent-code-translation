package main

import (
	"fmt"
	"sort"
)

func Median(l []int) float64 {
	// Create a copy to avoid modifying the original slice
	sorted := make([]int, len(l))
	copy(sorted, l)
	sort.Ints(sorted)

	n := len(sorted)
	if n%2 == 1 {
		// Odd length: return the middle element
		return float64(sorted[n/2])
	}
	// Even length: return the average of the two middle elements
	return (float64(sorted[n/2-1]) + float64(sorted[n/2])) / 2.0
}

func main() {
	fmt.Println(Median([]int{3, 1, 2, 4, 5}))
	fmt.Println(Median([]int{-10, 4, 6, 1000, 10, 20}))
}
