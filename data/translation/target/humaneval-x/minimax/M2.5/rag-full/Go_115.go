package main

import "math"

func MaxFill(grid [][]int, capacity int) int {
	total := 0
	for _, row := range grid {
		rowSum := 0
		for _, val := range row {
			rowSum += val
		}
		// Use integer ceiling division: (a + b - 1) / b
		// This avoids converting to float and using math.Ceil
		times := (rowSum + capacity - 1) / capacity
		total += times
	}
	return total
}