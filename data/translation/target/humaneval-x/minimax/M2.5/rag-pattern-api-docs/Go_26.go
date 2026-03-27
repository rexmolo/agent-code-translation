package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number
	counts := make(map[int]int)
	for _, n := range numbers {
		counts[n]++
	}

	// Filter to keep only elements that appear once (count <= 1)
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
	result := RemoveDuplicates([]int{1, 2, 3, 2, 4})
	fmt.Println(result) // Output: [1 3 4]
}
