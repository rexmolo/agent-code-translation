package main

import (
	"fmt"
	"sort"
)

// SortThird takes a slice l and returns a new slice l' such that
// l' is identical to l in the indicies that are not divisible by three, while its values at the indicies that are divisible by three are equal
// to the values of the corresponding indicies of l, but sorted.
func SortThird(l []int) []int {
	// Create a new slice to store the elements at indices divisible by 3
	thirdElements := []int{}
	for i, v := range l {
		if i%3 == 0 {
			thirdElements = append(thirdElements, v)
		}
	}

	// Sort these elements
	sort.Ints(thirdElements)

	// Create a copy of the original slice to modify and return.
	// This mimics the Python `l = list(l)` behavior which creates a shallow copy.
	result := make([]int, len(l))
	copy(result, l)

	// Place the sorted elements back into the result slice at the correct positions
	sortedIndex := 0
	for i := range result {
		if i%3 == 0 {
			result[i] = thirdElements[sortedIndex]
			sortedIndex++
		}
	}

	return result
}

// main function to demonstrate SortThird with the examples from the docstring.
func main() {
	fmt.Println(SortThird([]int{1, 2, 3}))
	// Expected: [1 2 3]

	fmt.Println(SortThird([]int{5, 6, 3, 4, 8, 9, 2}))
	// Expected: [2 6 3 4 8 9 5]
}
