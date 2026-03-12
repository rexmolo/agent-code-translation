package main

import (
	"fmt"
)

func Minpath(grid [][]int, k int) []int {
	n := len(grid)

	// Find the position of value 1 in the grid
	i1, j1 := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				i1 = i
				j1 = j
				break
			}
		}
	}

	// Find the minimum value among neighbors of the cell containing 1
	val := n*n + 1 // Initialize with a value larger than any possible neighbor
	temp := []int{}

	if i1 != 0 {
		temp = append(temp, grid[i1-1][j1])
	}
	if j1 != 0 {
		temp = append(temp, grid[i1][j1-1])
	}
	if i1 != n-1 {
		temp = append(temp, grid[i1+1][j1])
	}
	if j1 != n-1 {
		temp = append(temp, grid[i1][j1+1])
	}

	// Find minimum in temp slice
	val = temp[0]
	for _, v := range temp {
		if v < val {
			val = v
		}
	}

	// Build the answer: even indices get 1, odd indices get the minimum neighbor value
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
