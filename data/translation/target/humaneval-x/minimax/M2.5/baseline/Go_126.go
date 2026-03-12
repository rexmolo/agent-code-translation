package main

import "fmt"

func IsSorted(lst []int) bool {
	// Count occurrences of each number
	countDigit := make(map[int]int)
	for _, n := range lst {
		countDigit[n]++
	}

	// Check if any number appears more than 2 times
	for _, n := range lst {
		if countDigit[n] > 2 {
			return false
		}
	}

	// Check if list is sorted in ascending order
	for i := 1; i < len(lst); i++ {
		if lst[i-1] > lst[i] {
			return false
		}
	}

	return true
}

func main() {
	// Test cases
	fmt.Println(IsSorted([]int{5}))                      // true
	fmt.Println(IsSorted([]int{1, 2, 3, 4, 5}))          // true
	fmt.Println(IsSorted([]int{1, 3, 2, 4, 5}))          // false
	fmt.Println(IsSorted([]int{1, 2, 3, 4, 5, 6}))       // true
	fmt.Println(IsSorted([]int{1, 2, 3, 4, 5, 6, 7}))    // true
	fmt.Println(IsSorted([]int{1, 3, 2, 4, 5, 6, 7}))    // false
	fmt.Println(IsSorted([]int{1, 2, 2, 3, 3, 4}))       // true
	fmt.Println(IsSorted([]int{1, 2, 2, 2, 3, 4}))       // false
}
