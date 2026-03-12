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
		buckets := math.Ceil(float64(sum) / float64(capacity))
		total += int(buckets)
	}
	return total
}

func main() {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Read grid dimensions
	scanner.Scan()
	dims := strings.Fields(scanner.Text())
	if len(dims) < 2 {
		return
	}
	h, _ := strconv.Atoi(dims[0])
	w, _ := strconv.Atoi(dims[1])

	// Read capacity
	scanner.Scan()
	capacity, _ := strconv.Atoi(scanner.Text())

	// Read grid
	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		grid[i] = make([]int, w)
		for j := 0; j < w && j < len(row); j++ {
			grid[i][j], _ = strconv.Atoi(row[j])
		}
	}

	result := MaxFill(grid, capacity)
	fmt.Println(result)
}