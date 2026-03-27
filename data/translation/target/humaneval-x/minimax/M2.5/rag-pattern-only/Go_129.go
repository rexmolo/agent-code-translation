func Minpath(grid [][]int, k int) []int {
	n := len(grid)

	// Find position of value 1 and minimum neighbor value
	var minNeighbor int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// Initialize with max possible value
				minNeighbor = n * n

				// Check all four neighbors and find minimum
				if i > 0 && grid[i-1][j] < minNeighbor {
					minNeighbor = grid[i-1][j]
				}
				if i < n-1 && grid[i+1][j] < minNeighbor {
					minNeighbor = grid[i+1][j]
				}
				if j > 0 && grid[i][j-1] < minNeighbor {
					minNeighbor = grid[i][j-1]
				}
				if j < n-1 && grid[i][j+1] < minNeighbor {
					minNeighbor = grid[i][j+1]
				}
			}
		}
	}

	// Build answer: alternating 1 and minNeighbor
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		if i%2 == 0 {
			ans[i] = 1
		} else {
			ans[i] = minNeighbor
		}
	}

	return ans
}