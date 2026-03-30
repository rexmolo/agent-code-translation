package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number using a map
	count := make(map[int]int)
	for _, n := range numbers {
		count[n]++
	}

	// Filter numbers that appear exactly once (count <= 1)
	result := make([]int, 0, len(numbers))
	for _, n := range numbers {
		if count[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	// Example usage to verify the function works correctly
	result := RemoveDuplicates([]int{1, 2, 3, 2, 4})
	fmt.Println(result)
}
