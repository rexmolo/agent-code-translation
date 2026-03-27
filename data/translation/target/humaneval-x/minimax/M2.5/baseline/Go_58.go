package main

import (
	"fmt"
	"sort"
)

func Common(l1 []int, l2 []int) []int {
	// Use a map as a set to store unique common elements
	seen := make(map[int]bool)

	for _, e1 := range l1 {
		for _, e2 := range l2 {
			if e1 == e2 {
				seen[e1] = true
			}
		}
	}

	// Extract keys from map and sort them
	result := make([]int, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}

	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121}))
	fmt.Println(Common([]int{5, 3, 2, 8}, []int{3, 2}))
}