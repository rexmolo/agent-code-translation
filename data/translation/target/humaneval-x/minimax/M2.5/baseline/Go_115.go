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
		count := 0
		for _, val := range row {
			count += val
		}
		// Ceiling division: (count + capacity - 1) / capacity
		total += (count + capacity - 1) / capacity
	}
	return total
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter grid as JSON-like format [[0,0,1,0],[0,1,0,0]]:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")

	// Parse grid: remove brackets and split by '],['
	gridStr := strings.Trim(input, "[]")
	if gridStr == "" {
		fmt.Println(0)
		return
	}

	rowsStr := strings.Split(gridStr, "],[")
	var grid [][]int
	for _, rowStr := range rowsStr {
		rowStr = strings.Trim(rowStr, "[]")
		values := strings.Split(rowStr, ",")
		var row []int
		for _, v := range values {
			val, _ := strconv.Atoi(v)
			row = append(row, val)
		}
		grid = append(grid, row)
	}

	fmt.Println("Enter bucket capacity:")
	input, _ = reader.ReadString('\n')
	capacity, _ := strconv.Atoi(strings.TrimSpace(input))

	result := MaxFill(grid, capacity)
	fmt.Println(result)
}
