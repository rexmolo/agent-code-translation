package main

import (
	"math"
)

func MaxFill(grid [][]int, capacity int) int {
	total := 0
	for _, row := range grid {
		// Count the number of 1s (water units) in the row
		rowSum := 0
		for _, val := range row {
			rowSum += val
		}
		// Calculate ceiling of division: number of bucket drops needed for this row
		// Using math.Ceil to handle the rounding up
		drops := int(math.Ceil(float64(rowSum) / float64(capacity)))
		total += drops
	}
	return total
}
