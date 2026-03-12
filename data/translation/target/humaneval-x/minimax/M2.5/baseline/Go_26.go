package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number
	count := make(map[int]int)
	for _, n := range numbers {
		count[n]++
	}

	// Keep only numbers that appear once (count <= 1)
	result := []int{}
	for _, n := range numbers {
		if count[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	// Test the function
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 2, 4}))
}
