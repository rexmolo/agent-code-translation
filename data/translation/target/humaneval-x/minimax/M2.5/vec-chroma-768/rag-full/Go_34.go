package main

import "slices"

func Unique(l []int) []int {
	// Use map[int]struct{} as a set to track seen elements
	seen := make(map[int]struct{})
	result := make([]int, 0, len(l))

	for _, v := range l {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	slices.Sort(result)
	return result
}

// Test with main function
import "fmt"

func main() {
	test := []int{5, 3, 5, 2, 3, 3, 9, 0, 123}
	fmt.Println(Unique(test))
}
