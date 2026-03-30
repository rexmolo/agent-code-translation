package main

import (
	"fmt"
	"sort"
)

func SortArray(arr []int) []int {
	// Create a copy to avoid modifying the original
	result := make([]int, len(arr))
	copy(result, arr)

	// Sort by number of 1s in binary (ascending), then by value for ties
	sort.Slice(result, func(i, j int) bool {
		countI := countOnes(result[i])
		countJ := countOnes(result[j])
		if countI != countJ {
			return countI < countJ
		}
		return result[i] < result[j]
	})

	return result
}

// countOnes counts the number of 1 bits in the binary representation of n
func countOnes(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func main() {
	fmt.Println(SortArray([]int{1, 5, 2, 3, 4}))
	fmt.Println(SortArray([]int{-2, -3, -4, -5, -6}))
	fmt.Println(SortArray([]int{1, 0, 2, 3, 4}))
}
