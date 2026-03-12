package main

import "fmt"

// RemoveDuplicates removes all elements that occur more than once from a slice of integers.
// It keeps the order of the remaining elements the same as in the input.
func RemoveDuplicates(numbers []int) []int {
	// Create a map to store the frequency of each number.
	counts := make(map[int]int)
	for _, num := range numbers {
		counts[num]++
	}

	// Create a result slice. Using a nil slice and append is idiomatic.
	var result []int

	// Iterate through the original slice to maintain the order.
	for _, num := range numbers {
		// If a number appears exactly once, add it to the result.
		if counts[num] == 1 {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	// Example from the Python docstring
	input := []int{1, 2, 3, 2, 4}
	output := RemoveDuplicates(input)
	fmt.Println(output) // Expected output: [1 3 4]
}
