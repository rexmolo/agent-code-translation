package main

func Minpath(grid [][]int, k int) []int {
	// The logic to find the lexicographically smallest path is to always choose
	// the smallest available value at each step. The absolute smallest value in the
	// grid is 1. Therefore, the optimal path must start with 1.
	// The second step must be to a neighbor of 1. To keep the path lexicographically
	// small, we choose the neighbor with the minimum value.
	// For the third step, we are at the minimum-value neighbor. The best move is to
	// return to the cell with value 1, as it's the smallest possible value.
	// This establishes a pattern for the minimum path: alternating between 1 and its
	// smallest neighbor.

	// If the path length is 1, the answer is simply [1].
	if k == 1 {
		return []int{1}
	}

	n := len(grid)

	// Find the coordinates of the cell with value 1.
	var r1, c1 int
findLoop:
	for r, row := range grid {
		for c, val := range row {
			if val == 1 {
				r1 = r
				c1 = c
				break findLoop
			}
		}
	}

	// Find the minimum value among the neighbors of 1.
	// Initialize with a value guaranteed to be larger than any grid value.
	minNeighborVal := n*n + 1

	// Check up
	if r1 > 0 {
		if grid[r1-1][c1] < minNeighborVal {
			minNeighborVal = grid[r1-1][c1]
		}
	}
	// Check down
	if r1 < n-1 {
		if grid[r1+1][c1] < minNeighborVal {
			minNeighborVal = grid[r1+1][c1]
		}
	}
	// Check left
	if c1 > 0 {
		if grid[r1][c1-1] < minNeighborVal {
			minNeighborVal = grid[r1][c1-1]
		}
	}
	// Check right
	if c1 < n-1 {
		if grid[r1][c1+1] < minNeighborVal {
			minNeighborVal = grid[r1][c1+1]
		}
	}

	// Construct the result path by alternating between 1 and its smallest neighbor.
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		if i%2 == 0 {
			ans[i] = 1
		} else {
			ans[i] = minNeighborVal
		}
	}

	return ans
}
