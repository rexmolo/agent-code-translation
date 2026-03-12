package main

import (
	"fmt"
)

// Intersperse inserts a number 'delimeter' between every two consecutive elements of the input slice `numbers`.
func Intersperse(numbers []int, delimeter int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	// For a slice of length n > 0, the result will have n + (n-1) = 2n-1 elements.
	// Pre-allocating capacity is more efficient than repeated reallocations.
	result := make([]int, 0, 2*len(numbers)-1)

	for i, n := range numbers {
		if i > 0 {
			result = append(result, delimeter)
		}
		result = append(result, n)
	}
	return result
}

func main() {
	// Example 1 from Python docstring
	fmt.Println(Intersperse([]int{}, 4))

	// Example 2 from Python docstring
	fmt.Println(Intersperse([]int{1, 2, 3}, 4))

	// Additional test cases
	fmt.Println(Intersperse([]int{1}, 99))
	fmt.Println(Intersperse([]int{1, 2}, 0))
}
