package main

import (
	"fmt"
	"sort"
)

func Median(l []int) float64 {
	// Sort the slice in ascending order
	sort.Ints(l)

	n := len(l)
	if n%2 == 1 {
		// Odd length: return the middle element
		return float64(l[n/2])
	}
	// Even length: return the average of the two middle elements
	return (float64(l[n/2-1]) + float64(l[n/2])) / 2.0
}

func main() {
	// Test cases from docstrings
	fmt.Println(Median([]int{3, 1, 2, 4, 5}))           // Output: 3
	fmt.Println(Median([]int{-10, 4, 6, 1000, 10, 20})) // Output: 15
}
