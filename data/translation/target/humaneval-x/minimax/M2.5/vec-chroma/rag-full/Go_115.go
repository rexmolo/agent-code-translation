package main

import "math"

// MaxFill returns the number of times buckets need to be lowered to empty all wells.
// Each row in the grid represents a well, and each 1 represents a unit of water.
// All buckets have the same capacity.
func MaxFill(grid [][]int, capacity int) int {
	total := 0

	for _, row := range grid {
		rowSum := 0
		for _, val := range row {
			rowSum += val
		}
		// Calculate ceiling of rowSum divided by capacity
		drops := int(math.Ceil(float64(rowSum) / float64(capacity)))
		total += drops
	}

	return total
}
