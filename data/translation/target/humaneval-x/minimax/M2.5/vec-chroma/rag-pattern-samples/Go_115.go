package main

func MaxFill(grid [][]int, capacity int) int {
    total := 0
    for _, row := range grid {
        rowSum := 0
        for _, val := range row {
            rowSum += val
        }
        // Ceiling division: (a + b - 1) / b for positive integers
        drops := (rowSum + capacity - 1) / capacity
        total += drops
    }
    return total
}
