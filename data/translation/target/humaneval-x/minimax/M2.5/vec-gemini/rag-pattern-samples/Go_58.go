package main

import (
	"fmt"
	"sort"
)

func Common(l1 []int, l2 []int) []int {
	// Create a map to represent a set of elements from l1
	set := make(map[int]struct{})
	for _, e := range l1 {
		set[e] = struct{}{}
	}

	// Find common elements by checking if elements from l2 exist in the set
	var result []int
	for _, e := range l2 {
		if _, ok := set[e]; ok {
			result = append(result, e)
		}
	}

	// Remove duplicates using a map
	unique := make(map[int]struct{})
	for _, e := range result {
		unique[e] = struct{}{}
	}

	// Convert to slice and sort
	sortedResult := make([]int, 0, len(unique))
	for e := range unique {
		sortedResult = append(sortedResult, e)
	}

	sort.Ints(sortedResult)
	return sortedResult
}

func main() {
	// Test with the examples from the docstring
	result1 := Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121})
	fmt.Println(result1)

	result2 := Common([]int{5, 3, 2, 8}, []int{3, 2})
	fmt.Println(result2)
}
