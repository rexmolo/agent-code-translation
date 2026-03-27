package main

import (
	"fmt"
	"sort"
)

func Median(l []int) float64 {
	sorted := make([]int, len(l))
	copy(sorted, l)
	sort.Ints(sorted)

	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	} else {
		return float64(sorted[n/2-1]+sorted[n/2]) / 2.0
	}
}

func main() {
	// Test cases
	fmt.Println(Median([]int{3, 1, 2, 4, 5}))           // 3
	fmt.Println(Median([]int{-10, 4, 6, 1000, 10, 20})) // 15
}
