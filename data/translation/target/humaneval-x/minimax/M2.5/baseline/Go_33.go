package main

import (
	"fmt"
	"sort"
)

func SortThird(l []int) []int {
	// Extract elements at indices divisible by 3 (0, 3, 6, 9, ...)
	var third []int
	for i := 0; i < len(l); i += 3 {
		third = append(third, l[i])
	}

	// Sort the third elements
	sort.Ints(third)

	// Reconstruct the result: sorted values at indices divisible by 3,
	// original values elsewhere
	result := make([]int, len(l))
	thirdIndex := 0
	for i := 0; i < len(l); i++ {
		if i%3 == 0 {
			result[i] = third[thirdIndex]
			thirdIndex++
		} else {
			result[i] = l[i]
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(SortThird([]int{1, 2, 3}))          // [1 2 3]
	fmt.Println(SortThird([]int{5, 6, 3, 4, 8, 9, 2})) // [2 6 3 4 8 9 5]
}
