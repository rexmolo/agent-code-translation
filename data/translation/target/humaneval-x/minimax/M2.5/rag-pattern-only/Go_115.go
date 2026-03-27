package main

import "fmt"

func MaxFill(grid [][]int, capacity int) int {
	total := 0
	for _, row := range grid {
		sum := 0
		for _, val := range row {
			sum += val
		}
		// Calculate ceiling of sum/capacity using integer arithmetic
		// This avoids importing math package
		if sum > 0 {
			total += (sum + capacity - 1) / capacity
		}
		// When sum is 0, we add 0 (no bucket drops needed)
	}
	return total
}

func main() {
	// Test cases
	grid1 := [][]int{{0, 0, 1, 0}, {0, 1, 0, 0}, {1, 1, 1, 1}}
	fmt.Println(MaxFill(grid1, 1)) // Expected: 6

	grid2 := [][]int{{0, 0, 1, 1}, {0, 0, 0, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}}
	fmt.Println(MaxFill(grid2, 2)) // Expected: 5

	grid3 := [][]int{{0, 0, 0}, {0, 0, 0}}
	fmt.Println(MaxFill(grid3, 5)) // Expected: 0
}
