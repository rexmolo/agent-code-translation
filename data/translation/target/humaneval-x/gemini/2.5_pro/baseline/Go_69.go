package main

import "fmt"

// Search is given a non-empty list of positive integers. It returns the greatest integer that is greater than
// zero, and has a frequency greater than or equal to the value of the integer itself.
// The frequency of an integer is the number of times it appears in the list.
// If no such a value exists, it returns -1.
// Examples:
//     Search([]int{4, 1, 2, 2, 3, 1}) == 2
//     Search([]int{1, 2, 2, 3, 3, 3, 4, 4, 4}) == 3
//     Search([]int{5, 5, 4, 4, 4}) == -1
func Search(lst []int) int {
	// The problem description implies a non-empty list, but handling this
	// edge case makes the function more robust.
	if len(lst) == 0 {
		return -1
	}

	// Use a map to count the frequency of each number. This is idiomatic in Go
	// and more memory-efficient than the original Python's slice-based approach
	// if the input numbers are large but sparse.
	freq := make(map[int]int)
	for _, num := range lst {
		freq[num]++
	}

	ans := -1
	// Iterate through the frequency map to find the largest qualifying number.
	for num, count := range freq {
		// A qualifying number must be positive (greater than zero), and its frequency
		// must be greater than or equal to its value.
		if num > 0 && count >= num {
			// If the current qualifying number is greater than our best answer so far,
			// update the answer. This is necessary because map iteration order is not guaranteed.
			if num > ans {
				ans = num
			}
		}
	}

	return ans
}

// main function to serve as an entry point and demonstrate the Search function.
func main() {
	fmt.Println(Search([]int{4, 1, 2, 2, 3, 1}))
	fmt.Println(Search([]int{1, 2, 2, 3, 3, 3, 4, 4, 4}))
	fmt.Println(Search([]int{5, 5, 4, 4, 4}))
}
