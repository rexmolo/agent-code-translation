package main

import (
    "math"
)

func MaxFill(grid [][]int, capacity int) int {
    total := 0
    for _, row := range grid {
        count := 0
        for _, val := range row {
            count += val
        }
        // Calculate ceiling of count/capacity
        drops := math.Ceil(float64(count) / float64(capacity))
        total += int(drops)
    }
    return total
}

func main() {
    // Test examples
    grid1 := [][]int{{0, 0, 1, 0}, {0, 1, 0, 0}, {1, 1, 1, 1}}
    grid2 := [][]int{{0, 0, 1, 1}, {0, 0, 0, 0}, {1, 1, 1, 1}, {0, 1, 1, 1}}
    grid3 := [][]int{{0, 0, 0}, {0, 0, 0}}

    println(MaxFill(grid1, 1)) // Expected: 6
    println(MaxFill(grid2, 2)) // Expected: 5
    println(MaxFill(grid3, 5)) // Expected: 0
}
