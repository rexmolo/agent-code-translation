package main

import (
	"fmt"
	"sort"
)

// GetRow finds all occurrences of an integer x in a 2D slice and returns their coordinates.
// The coordinates are sorted first by row (ascending) and then by column (descending).
func GetRow(lst [][]int, x int) [][2]int {
	var coords [][2]int

	// Find all coordinates where the value is x
	for i, row := range lst {
		for j, val := range row {
			if val == x {
				coords = append(coords, [2]int{i, j})
			}
		}
	}

	// Sort the coordinates
	// The Python code `sorted(sorted(coords, key=lambda x: x[1], reverse=True), key=lambda x: x[0])`
	// is a stable sort trick. The final sort by row is the primary key.
	// We can achieve the same result in Go with a single sort operation with a custom comparison.
	sort.Slice(coords, func(i, j int) bool {
		// Primary sort key: row index, ascending
		if coords[i][0] != coords[j][0] {
			return coords[i][0] < coords[j][0]
		}
		// Secondary sort key: column index, descending
		return coords[i][1] > coords[j][1]
	})

	return coords
}

// main function to run the examples from the Python docstring.
func main() {
	lst1 := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 1, 6},
		{1, 2, 3, 4, 5, 1},
	}
	fmt.Printf("Input: %v, %d\n", lst1, 1)
	// Expected: [(0, 0), (1, 4), (1, 0), (2, 5), (2, 0)]
	// Go output: [[0 0] [1 4] [1 0] [2 5] [2 0]]
	fmt.Printf("Result: %v\n\n", GetRow(lst1, 1))

	lst2 := [][]int{}
	fmt.Printf("Input: %v, %d\n", lst2, 1)
	// Expected: []
	// Go output: []
	fmt.Printf("Result: %v\n\n", GetRow(lst2, 1))

	lst3 := [][]int{{}, {1}, {1, 2, 3}}
	fmt.Printf("Input: %v, %d\n", lst3, 3)
	// Expected: [(2, 2)]
	// Go output: [[2 2]]
	fmt.Printf("Result: %v\n\n", GetRow(lst3, 3))
}
