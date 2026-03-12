package main

import "fmt"

func Solution(lst []int) int {
	sum := 0
	for idx, x := range lst {
		if idx%2 == 0 && x%2 == 1 {
			sum += x
		}
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println(Solution([]int{5, 8, 7, 1}))        // Expected: 12
	fmt.Println(Solution([]int{3, 3, 3, 3, 3}))   // Expected: 9
	fmt.Println(Solution([]int{30, 13, 24, 321})) // Expected: 0
}
