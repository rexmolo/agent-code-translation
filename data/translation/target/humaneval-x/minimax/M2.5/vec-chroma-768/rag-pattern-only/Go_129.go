package main

func Minpath(grid [][]int, k int) []int {
	n := len(grid)
	val := n*n + 1

	// Find the cell containing 1 and get the minimum neighbor value
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// Check all four neighbors (up, down, left, right)
				if i > 0 && grid[i-1][j] < val {
					val = grid[i-1][j]
				}
				if j > 0 && grid[i][j-1] < val {
					val = grid[i][j-1]
				}
				if i < n-1 && grid[i+1][j] < val {
					val = grid[i+1][j]
				}
				if j < n-1 && grid[i][j+1] < val {
					val = grid[i][j+1]
				}
			}
		}
	}

	// Build the result: even indices get 1, odd indices get val
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		if i%2 == 0 {
			ans[i] = 1
		} else {
			ans[i] = val
		}
	}
	return ans
}
