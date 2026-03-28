package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number using a map (equivalent to Counter)
	count := make(map[int]int)
	for _, n := range numbers {
		count[n]++
	}

	// Filter elements that appear only once, preserving order
	var result []int
	for _, n := range numbers {
		if count[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	// Test the function
	result := RemoveDuplicates([]int{1, 2, 3, 2, 4})
	fmt.Println(result)
}