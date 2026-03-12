package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func Minpath(grid [][]int, k int) []int {
    n := len(grid)
    val := n*n + 1
    
    // Find the position of cell with value 1
    var iPos, jPos int
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                iPos = i
                jPos = j
                break
            }
        }
    }
    
    // Find minimum value among neighbors of cell with value 1
    temp := []int{}
    if iPos != 0 {
        temp = append(temp, grid[iPos-1][jPos])
    }
    if jPos != 0 {
        temp = append(temp, grid[iPos][jPos-1])
    }
    if iPos != n-1 {
        temp = append(temp, grid[iPos+1][jPos])
    }
    if jPos != n-1 {
        temp = append(temp, grid[iPos][jPos+1])
    }
    
    val = temp[0]
    for _, v := range temp {
        if v < val {
            val = v
        }
    }
    
    // Build the answer path
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
    // Read input
    reader := bufio.NewReader(os.Stdin)
    
    var n int
    fmt.Scanf("%d", &n)
    
    grid := make([][]int, n)
    for i := 0; i < n; i++ {
        grid[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scanf("%d", &grid[i][j])
        }
    }
    
    var k int
    fmt.Scanf("%d", &k)
    
    result := Minpath(grid, k)
    
    // Output as JSON array
    jsonBytes, _ := json.Marshal(result)
    fmt.Println(string(jsonBytes))
}
