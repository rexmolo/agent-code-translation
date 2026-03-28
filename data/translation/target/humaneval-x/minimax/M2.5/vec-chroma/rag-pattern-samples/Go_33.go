package main

import (
	"fmt"
	"sort"
)

func main() {
	// Example usage
	fmt.Println(sortThird([]int{1, 2, 3}))
	fmt.Println(sortThird([]int{5, 6, 3, 4, 8, 9, 2}))
}

func SortThird(l []int) []int {
	// Extract elements at indices divisible by 3 (0, 3, 6, ...)
	var thirdElements []int
	for i := 0; i < len(l); i += 3 {
		thirdElements = append(thirdElements, l[i])
	}

	// Sort the extracted elements
	sort.Ints(thirdElements)

	// Place sorted elements back into the original positions
	for i := 0; i < len(l); i += 3 {
		l[i] = thirdElements[i/3]
	}

	return l
}