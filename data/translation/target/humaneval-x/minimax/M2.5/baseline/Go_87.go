package main

import (
	"fmt"
	"sort"
)

func GetRow(lst [][]int, x int) [][2]int {
	var coords [][2]int

	// Find all coordinates where lst[i][j] == x
	for i := 0; i < len(lst); i++ {
		for j := 0; j < len(lst[i]); j++ {
			if lst[i][j] == x {
				coords = append(coords, [2]int{i, j})
			}
		}
	}

	// Sort: first by column descending, then by row ascending (stable sort)
	// This matches Python's: sorted(sorted(coords, key=lambda x: x[1], reverse=True), key=lambda x: x[0])
	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i][1] > coords[j][1] // Sort by column descending
	})

	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i][0] < coords[j][0] // Sort by row ascending
	})

	return coords
}

func main() {
	// Test cases
	fmt.Println(GetRow([][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 1, 6},
		{1, 2, 3, 4, 5, 1},
	}, 1))
	// Output: [[0 0] [1 4] [1 0] [2 5] [2 0]]

	fmt.Println(GetRow([][]int{}, 1))
	// Output: []

	fmt.Println(GetRow([][]int{
		{},
		{1},
		{1, 2, 3},
	}, 3))
	// Output: [[2 2]]
}
