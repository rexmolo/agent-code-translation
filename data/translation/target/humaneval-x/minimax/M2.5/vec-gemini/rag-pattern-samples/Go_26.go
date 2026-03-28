package main

func RemoveDuplicates(numbers []int) []int {
	// Count occurrences of each number
	count := make(map[int]int)
	for _, n := range numbers {
		count[n]++
	}

	// Filter numbers that appear only once (count <= 1)
	var result []int
	for _, n := range numbers {
		if count[n] <= 1 {
			result = append(result, n)
		}
	}

	return result
}