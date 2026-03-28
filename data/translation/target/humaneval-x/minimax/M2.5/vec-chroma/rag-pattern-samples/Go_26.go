package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number using a map
	counts := make(map[int]int)
	for _, n := range numbers {
		counts[n]++
	}

	// Filter to keep only elements that appear exactly once
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
	nums := []int{1, 2, 3, 2, 4}
	result := RemoveDuplicates(nums)
	fmt.Println(result)
}
