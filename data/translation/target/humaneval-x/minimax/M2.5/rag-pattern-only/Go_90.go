package main

import (
	"fmt"
	"sort"
)

func NextSmallest(lst []int) interface{} {
	if len(lst) < 2 {
		return nil
	}

	// Remove duplicates using a map
	unique := make(map[int]bool)
	for _, v := range lst {
		unique[v] = true
	}

	// Convert map keys to slice
	uniqueList := make([]int, 0, len(unique))
	for k := range unique {
		uniqueList = append(uniqueList, k)
	}

	// Sort the slice
	sort.Ints(uniqueList)

	// Return nil if less than 2 unique elements, otherwise return 2nd smallest
	if len(uniqueList) < 2 {
		return nil
	}

	return uniqueList[1]
}

func main() {
	// Test cases
	fmt.Println(NextSmallest([]int{1, 2, 3, 4, 5})) // 2
	fmt.Println(NextSmallest([]int{5, 1, 4, 3, 2})) // 2
	fmt.Println(NextSmallest([]int{}))              // nil
	fmt.Println(NextSmallest([]int{1, 1}))          // nil
}
