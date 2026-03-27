package main

import "bufio"
import "fmt"
import "os"

func MaxFill(grid [][]int, capacity int) int {
    total := 0
    for _, row := range grid {
        rowSum := 0
        for _, val := range row {
            rowSum += val
        }
        // Ceiling division: (a + b - 1) / b
        total += (rowSum + capacity - 1) / capacity
    }
    return total
}

func main() {
    // Example usage based on the problem description
    // Read input from stdin
    reader := bufio.NewReader(os.Stdin)

    // Read grid dimensions and capacity
    var rows, cols, capacity int
    fmt.Fscan(reader, &rows, &cols, &capacity)

    grid := make([][]int, rows)
    for i := 0; i < rows; i++ {
        grid[i] = make([]int, cols)
        for j := 0; j < cols; j++ {
            fmt.Fscan(reader, &grid[i][j])
        }
    }

    result := MaxFill(grid, capacity)
    fmt.Println(result)
}
