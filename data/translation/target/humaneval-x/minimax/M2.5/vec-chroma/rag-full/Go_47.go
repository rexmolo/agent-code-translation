package main

import (
	"fmt"
	"sort"
)

func Median(l []int) float64 {
	// Make a copy to avoid modifying the original slice (preserves Python behavior)
	sorted := make([]int, len(l))
	copy(sorted, l)
	sort.Ints(sorted)

	n := len(sorted)
	if n%2 == 1 {
		// Odd length: return the middle element
		return float64(sorted[n/2])
	} else {
		// Even length: return the average of the two middle elements
		return float64(sorted[n/2-1]+sorted[n/2]) / 2.0
	}
}

func main() {
	// Test examples from Python docstring
	fmt.Println(Median([]int{3, 1, 2, 4, 5}))           // Expected: 3
	fmt.Println(Median([]int{-10, 4, 6, 1000, 10, 20})) // Expected: 15
}