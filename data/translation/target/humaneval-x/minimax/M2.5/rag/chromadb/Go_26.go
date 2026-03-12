package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number (equivalent to collections.Counter)
	count := make(map[int]int)
	for _, n := range numbers {
		count[n]++
	}

	// Filter: keep only elements that appear once (count <= 1)
	result := make([]int, 0)
	for _, n := range numbers {
		if count[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	// Test the function
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 2, 4})) // Output: [1 3 4]
}