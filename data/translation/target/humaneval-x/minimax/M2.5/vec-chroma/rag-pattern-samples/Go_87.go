package main

import (
    "fmt"
    "slices"
)

func GetRow(lst [][]int, x int) [][2]int {
    var coords [][2]int

    // Find all coordinates where lst[i][j] == x
    for i := 0; i < len(lst); i++ {
        for j := 0; j < len(lst[i]); j++ {
            if lst[i][j] == x {
                coords = append(coords, [2]int{i, j})
            }
        }
    }

    // Sort by row ascending first, then by column descending
    // Equivalent to Python's:
    // sorted(sorted(coords, key=lambda x: x[1], reverse=True), key=lambda x: x[0])
    slices.SortFunc(coords, func(a, b [2]int) int {
        if a[0] != b[0] {
            // Row ascending
            if a[0] < b[0] {
                return -1
            }
            return 1
        }
        // Row equal, sort by column descending
        if a[1] > b[1] {
            return -1
        }
        if a[1] < b[1] {
            return 1
        }
        return 0
    })

    return coords
}

func main() {
    // Test examples
    result1 := GetRow([][]int{
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 1, 6},
        {1, 2, 3, 4, 5, 1},
    }, 1)
    fmt.Printf("Test 1: %v\n", result1)
    // Expected: [[0 0] [1 4] [1 0] [2 5] [2 0]]

    result2 := GetRow([][]int{}, 1)
    fmt.Printf("Test 2: %v\n", result2)
    // Expected: []

    result3 := GetRow([][]int{{}, {1}, {1, 2, 3}}, 3)
    fmt.Printf("Test 3: %v\n", result3)
    // Expected: [[2 2]]
}