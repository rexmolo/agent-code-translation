package main

import (
	"bufio"
	"fmt"
	"math"
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
		if sum > 0 {
			bucketLowers := math.Ceil(float64(sum) / float64(capacity))
			total += int(bucketLowers)
		}
	}
	return total
}

func main() {
	// Read input from stdin
	// Format: first line is capacity, then grid rows as space-separated integers
	scanner := bufio.NewScanner(os.Stdin)
	
	var grid [][]int
	var capacity int
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		
		parts := strings.Fields(line)
		if len(parts) == 1 && capacity == 0 {
			// First non-empty line is capacity
			capacity, _ = strconv.Atoi(parts[0])
			continue
		}
		
		// Parse grid row
		var row []int
		for _, p := range parts {
			val, _ := strconv.Atoi(p)
			row = append(row, val)
		}
		if len(row) > 0 {
			grid = append(grid, row)
		}
	}

	result := MaxFill(grid, capacity)
	fmt.Println(result)
}