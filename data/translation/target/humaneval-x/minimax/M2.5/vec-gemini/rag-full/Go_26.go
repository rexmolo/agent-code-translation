package main

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number
	counts := make(map[int]int)
	for _, n := range numbers {
		counts[n]++
	}

	// Filter elements that appear only once (count <= 1)
	result := make([]int, 0, len(numbers))
	for _, n := range numbers {
		if counts[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}