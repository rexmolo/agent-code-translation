package main

import "slices"

func Minpath(grid [][]int, k int) []int {
    n := len(grid)
    val := n*n + 1 // Initialize to value larger than any possible answer

    // Find the cell containing value 1 and compute minimum of its neighbors
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                var temp []int
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
                val = slices.Min(temp)
            }
        }
    }

    // Build the answer: even indices get 1, odd indices get val
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
