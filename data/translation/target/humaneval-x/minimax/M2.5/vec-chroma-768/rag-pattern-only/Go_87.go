package main

import (
	"fmt"
	"slices"
)

func GetRow(lst [][]int, x int) [][2]int {
	var coords [][2]int

	// Find all coordinates where lst[i][j] == x
	for i := range lst {
		for j := range lst[i] {
			if lst[i][j] == x {
				coords = append(coords, [2]int{i, j})
			}
		}
	}

	// Sort by row ascending, then by column descending within each row
	slices.SortFunc(coords, func(a, b [2]int) int {
		// First sort by row ascending
		if a[0] != b[0] {
			if a[0] < b[0] {
				return -1
			}
			return 1
		}
		// If rows are equal, sort by column descending
		if b[1] != a[1] {
			if b[1] > a[1] {
				return -1
			}
			return 1
		}
		return 0
	})

	return coords
}

func main() {
	// Test case 1
	result1 := GetRow([][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 1, 6},
		{1, 2, 3, 4, 5, 1},
	}, 1)
	expected1 := [][2]int{{0, 0}, {1, 4}, {1, 0}, {2, 5}, {2, 0}}
	fmt.Println("Test 1:", result1)
	fmt.Println("Expected:", expected1)
	fmt.Println("Match:", fmt.Sprint(result1) == fmt.Sprint(expected1))

	// Test case 2
	result2 := GetRow([][]int{}, 1)
	expected2 := [][2]int{}
	fmt.Println("Test 2:", result2)
	fmt.Println("Expected:", expected2)
	fmt.Println("Match:", fmt.Sprint(result2) == fmt.Sprint(expected2))

	// Test case 3
	result3 := GetRow([][]int{{}, {1}, {1, 2, 3}}, 3)
	expected3 := [][2]int{{2, 2}}
	fmt.Println("Test 3:", result3)
	fmt.Println("Expected:", expected3)
	fmt.Println("Match:", fmt.Sprint(result3) == fmt.Sprint(expected3))
}