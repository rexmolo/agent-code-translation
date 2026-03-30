package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number
	counts := make(map[int]int)
	for _, n := range numbers {
		counts[n]++
	}

	// Filter numbers that appear exactly once (count <= 1)
	result := make([]int, 0, len(numbers))
	for _, n := range numbers {
		if counts[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	// Test the function
	numbers := []int{1, 2, 3, 2, 4}
	result := RemoveDuplicates(numbers)
	fmt.Println(result) // Output: [1 3 4]
}
