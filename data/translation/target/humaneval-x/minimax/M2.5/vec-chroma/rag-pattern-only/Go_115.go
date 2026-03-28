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
		rowSum := 0
		for _, val := range row {
			rowSum += val
		}
		// Calculate ceiling of rowSum/capacity using integer arithmetic
		times := rowSum / capacity
		if rowSum%capacity != 0 {
			times++
		}
		total += times
	}
	return total
}

func main() {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	
	var grid [][]int
	var capacity int
	
	fmt.Println("Enter grid rows (space-separated, 0 or 1), one per line.")
	fmt.Println("Enter empty line to finish grid input.")
	
	// Read grid rows
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		parts := strings.Fields(line)
		var row []int
		for _, p := range parts {
			v, _ := strconv.Atoi(p)
			row = append(row, v)
		}
		grid = append(grid, row)
	}
	
	fmt.Println("Enter bucket capacity:")
	scanner.Scan()
	capacity, _ = strconv.Atoi(strings.TrimSpace(scanner.Text()))
	
	result := MaxFill(grid, capacity)
	fmt.Println(result)
}
