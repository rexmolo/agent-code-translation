package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number using a map
	counts := make(map[int]int)
	for _, n := range numbers {
		counts[n]++
	}

	// Filter to keep only elements that occur once (count <= 1)
	var result []int
	for _, n := range numbers {
		if counts[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	// Example usage
	nums := []int{1, 2, 3, 2, 4}
	result := RemoveDuplicates(nums)
	fmt.Println(result) // Output: [1 3 4]
}
