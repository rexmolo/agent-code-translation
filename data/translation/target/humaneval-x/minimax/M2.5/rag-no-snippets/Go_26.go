package main

import "fmt"

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number using a map
	count := make(map[int]int)
	for _, n := range numbers {
		count[n]++
	}

	// Keep only elements that appear once (count <= 1)
	result := make([]int, 0)
	for _, n := range numbers {
		if count[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 2, 4}))
}