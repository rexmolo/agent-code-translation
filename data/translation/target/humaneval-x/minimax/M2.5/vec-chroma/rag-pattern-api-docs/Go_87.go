package main

import "sort"

func GetRow(lst [][]int, x int) [][2]int {
	var coords [][2]int

	// Build coordinates list (replaces Python list comprehension)
	for i := 0; i < len(lst); i++ {
		for j := 0; j < len(lst[i]); j++ {
			if lst[i][j] == x {
				coords = append(coords, [2]int{i, j})
			}
		}
	}

	// Python: sorted(sorted(coords, key=lambda x: x[1], reverse=True), key=lambda x: x[0])
	// First sort by column descending (reverse=True), then by row ascending
	// Use SliceStable to match Python's stable sort behavior
	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i][1] > coords[j][1]
	})

	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i][0] < coords[j][0]
	})

	return coords
}
