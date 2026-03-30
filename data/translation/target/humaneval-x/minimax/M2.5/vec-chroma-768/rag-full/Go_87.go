package main

import (
    "sort"
)

func GetRow(lst [][]int, x int) [][2]int {
    var coords [][2]int

    // Collect all coordinates where lst[i][j] == x
    // Equivalent to: [(i, j) for i in range(len(lst)) for j in range(len(lst[i])) if lst[i][j] == x]
    for i := 0; i < len(lst); i++ {
        for j := 0; j < len(lst[i]); j++ {
            if lst[i][j] == x {
                coords = append(coords, [2]int{i, j})
            }
        }
    }

    // Sort by row ascending, then by column descending
    // Equivalent to: sorted(sorted(coords, key=lambda x: x[1], reverse=True), key=lambda x: x[0])
    sort.Slice(coords, func(i, j int) bool {
        if coords[i][0] != coords[j][0] {
            return coords[i][0] < coords[j][0] // row ascending
        }
        return coords[i][1] > coords[j][1] // column descending within same row
    })

    return coords
}
