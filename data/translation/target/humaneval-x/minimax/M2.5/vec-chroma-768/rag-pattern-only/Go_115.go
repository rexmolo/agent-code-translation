package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MaxFill(grid [][]int, capacity int) int {
	total := 0
	for _, row := range grid {
		sum := 0
		for _, val := range row {
			sum += val
		}
		// Ceiling division: (sum + capacity - 1) / capacity
		drops := (sum + capacity - 1) / capacity
		total += drops
	}
	return total
}

func main() {
	// Read input from stdin for testing purposes
	// This is a simple test harness - adjust based on actual input format
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter grid as JSON-like format (e.g., [[0,0,1,0],[0,1,0,0]]): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	fmt.Print("Enter capacity: ")
	capacityInput, _ := reader.ReadString('\n')
	capacityInput = strings.TrimSpace(capacityInput)
	capacity, _ := strconv.Atoi(capacityInput)
	
	// Parse the grid (simplified parsing for demonstration)
	// For actual use, grid would be passed directly as function argument
	_ = input
	
	// Example usage:
	// grid := [][]int{{0, 0, 1, 0}, {0, 1, 0, 0}, {1, 1, 1, 1}}
	// result := MaxFill(grid, capacity)
	// fmt.Println(result)
	
	_ = capacity
}