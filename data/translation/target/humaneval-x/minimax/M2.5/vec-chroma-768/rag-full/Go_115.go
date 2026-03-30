package main

import (
	"math"
)

func MaxFill(grid [][]int, capacity int) int {
	total := 0
	for _, row := range grid {
		rowSum := 0
		for _, val := range row {
			rowSum += val
		}
		// Calculate ceil(rowSum / capacity) for each row
		if rowSum > 0 {
			total += int(math.Ceil(float64(rowSum) / float64(capacity)))
		}
	}
	return total
}