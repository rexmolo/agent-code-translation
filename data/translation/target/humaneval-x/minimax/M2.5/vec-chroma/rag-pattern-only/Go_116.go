package main

import (
	"fmt"
	"sort"
)

// countOnes counts the number of 1 bits in the binary representation of n
func countOnes(n int) int {
	count := 0
	if n < 0 {
		n = -n // Use absolute value for negative numbers to match Python's bin() behavior
	}
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func SortArray(arr []int) []int {
	// First sort by decimal value (equivalent to sorted(arr) in Python)
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	// Then stable sort by number of 1s in binary (equivalent to sorted(sorted(arr), key=...))
	// Using SliceStable to match Python's stable sort behavior
	sort.SliceStable(sortedArr, func(i, j int) int {
		countI := countOnes(sortedArr[i])
		countJ := countOnes(sortedArr[j])
		if countI != countJ {
			return countI - countJ
		}
		// If same number of 1s, sort by decimal value
		return sortedArr[i] - sortedArr[j]
	})

	return sortedArr
}

func main() {
	fmt.Println(SortArray([]int{1, 5, 2, 3, 4}))
	fmt.Println(SortArray([]int{-2, -3, -4, -5, -6}))
	fmt.Println(SortArray([]int{1, 0, 2, 3, 4}))
}