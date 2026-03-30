package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number using a map
	counts := make(map[int]int)
	for _, n := range numbers {
		counts[n]++
	}

	// Filter elements that appear only once, preserving order
	var result []int
	for _, n := range numbers {
		if counts[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	// Test the function
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 2, 4})) // Output: [1 3 4]
}