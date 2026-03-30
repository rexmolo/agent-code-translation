package main

import "sort"

func GetRow(lst [][]int, x int) [][2]int {
	var coords [][2]int

	// Find all coordinates where value equals x
	for i := 0; i < len(lst); i++ {
		for j := 0; j < len(lst[i]); j++ {
			if lst[i][j] == x {
				coords = append(coords, [2]int{i, j})
			}
		}
	}

	// Sort by column descending first
	sort.Slice(coords, func(i, j int) bool {
		return coords[i][1] > coords[j][1]
	})

	// Then sort by row ascending (stable to preserve column order within same row)
	sort.SliceStable(coords, func(i, j int) bool {
		return coords[i][0] < coords[j][0]
	})

	return coords
}
