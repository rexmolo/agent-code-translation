package main

import (
	"fmt"
	"sort"
)

func SortThird(l []int) []int {
	// Make a copy of the slice to avoid modifying the original
	result := make([]int, len(l))
	copy(result, l)

	// Collect elements at indices divisible by 3 (0, 3, 6, 9, ...)
	var third []int
	for i := 0; i < len(l); i += 3 {
		third = append(third, l[i])
	}

	// Sort the collected third elements
	sort.Ints(third)

	// Place sorted elements back at indices divisible by 3
	for i := 0; i < len(result); i += 3 {
		result[i] = third[i/3]
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(SortThird([]int{1, 2, 3}))           // [1 2 3]
	fmt.Println(SortThird([]int{5, 6, 3, 4, 8, 9, 2})) // [2 6 3 4 8 9 5]
}
