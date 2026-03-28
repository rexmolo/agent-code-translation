package main

import (
    "fmt"
    "sort"
)

func GetRow(lst [][]int, x int) [][2]int {
    coords := make([][2]int, 0)

    // Find all coordinates where lst[i][j] == x
    for i := range lst {
        for j := range lst[i] {
            if lst[i][j] == x {
                coords = append(coords, [2]int{i, j})
            }
        }
    }

    // Sort: first by row ascending, then by column descending
    // This mimics Python's: sorted(sorted(coords, key=lambda x: x[1], reverse=True), key=lambda x: x[0])
    sort.Slice(coords, func(i, j int) bool {
        if coords[i][0] != coords[j][0] {
            return coords[i][0] < coords[j][0] // row ascending
        }
        return coords[i][1] > coords[j][1] // column descending
    })

    return coords
}

func main() {
    // Test case 1
    result1 := GetRow([][]int{
        {1, 2, 3, 4, 5, 6},
        {1, 2, 3, 4, 1, 6},
        {1, 2, 3, 4, 5, 1},
    }, 1)
    fmt.Printf("%v\n", result1)

    // Test case 2
    result2 := GetRow([][]int{}, 1)
    fmt.Printf("%v\n", result2)

    // Test case 3
    result3 := GetRow([][]int{{}, {1}, {1, 2, 3}}, 3)
    fmt.Printf("%v\n", result3)
}