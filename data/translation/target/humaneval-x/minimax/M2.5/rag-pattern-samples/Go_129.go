package main

import "fmt"

func Minpath(grid [][]int, k int) []int {
	n := len(grid)
	val := n*n + 1

	// Find the cell with value 1 and get the minimum neighbor value
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				temp := make([]int, 0)
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
				val = temp[0]
				for _, v := range temp {
					if v < val {
						val = v
					}
				}
			}
		}
	}

	// Build answer: 1 at even indices, val at odd indices
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

func main() {
	// Test examples
	grid1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(Minpath(grid1, 3)) // Output: [1 2 1]

	grid2 := [][]int{{5, 9, 3}, {4, 1, 6}, {7, 8, 2}}
	fmt.Println(Minpath(grid2, 1)) // Output: [1]
}
