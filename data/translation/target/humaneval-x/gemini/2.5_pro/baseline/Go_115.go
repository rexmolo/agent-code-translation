package main

// MaxFill calculates the total number of times buckets need to be lowered to empty all wells.
func MaxFill(grid [][]int, capacity int) int {
	// totalFills accumulates the number of bucket lowerings for all wells.
	totalFills := 0

	// Iterate over each well (row) in the grid.
	for _, row := range grid {
		// Calculate the total amount of water in the current well.
		wellWater := 0
		for _, waterUnit := range row {
			wellWater += waterUnit
		}

		// If there is water in the well, calculate the fills needed.
		if wellWater > 0 {
			// The Python code uses `math.ceil(sum/capacity)`. This is a ceiling division.
			// In Go, an idiomatic and efficient way to perform ceiling division with positive
			// integers is to use the formula `(a + b - 1) / b`.
			fillsForRow := (wellWater + capacity - 1) / capacity
			totalFills += fillsForRow
		}
	}

	return totalFills
}
