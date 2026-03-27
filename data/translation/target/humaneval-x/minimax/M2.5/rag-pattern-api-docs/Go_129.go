package main

func Minpath(grid [][]int, k int) []int {
	n := len(grid)
	val := n*n + 1

	// Find the cell with value 1 and find the minimum neighbor value
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				temp := []int{}

				if i != 0 {
					temp = append(temp, grid[i-1][j])
				}

				if j != 0 {
					temp = append(temp, grid[i][j-1])
				}

				if i != n-1 {
					temp = append(temp, grid[i+1][j])
				}

				if j != n-1 {
					temp = append(temp, grid[i][j+1])
				}

				// Find minimum in temp slice
				minVal := temp[0]
				for _, v := range temp {
					if v < minVal {
						minVal = v
					}
				}
				val = minVal
			}
		}
	}

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
