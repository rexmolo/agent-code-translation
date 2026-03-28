package main

import (
    "fmt"
)

func Minpath(grid [][]int, k int) []int {
    n := len(grid)
    val := n*n + 1
    
    // Find the cell with value 1 and its neighbors
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
                
                // Find minimum in temp
                val = temp[0]
                for _, v := range temp {
                    if v < val {
                        val = v
                    }
                }
            }
        }
    }
    
    // Build answer array
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
    // Test example 1
    grid1 := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    result1 := Minpath(grid1, 3)
    fmt.Println(result1) // Expected: [1 2 1]
    
    // Test example 2
    grid2 := [][]int{
        {5, 9, 3},
        {4, 1, 6},
        {7, 8, 2},
    }
    result2 := Minpath(grid2, 1)
    fmt.Println(result2) // Expected: [1]
}
