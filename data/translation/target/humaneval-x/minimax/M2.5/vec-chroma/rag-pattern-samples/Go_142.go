package main

import "fmt"

func SumSquares(lst []int) int {
	result := 0
	for i, val := range lst {
		if i%3 == 0 {
			result += val * val
		} else if i%4 == 0 && i%3 != 0 {
			result += val * val * val
		} else {
			result += val
		}
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(SumSquares([]int{1, 2, 3}))           // Output: 6
	fmt.Println(SumSquares([]int{}))                  // Output: 0
	fmt.Println(SumSquares([]int{-1, -5, 2, -1, -5})) // Output: -126
}