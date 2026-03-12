package main

import (
	"sort"
)

func SortArray(arr []int) []int {
	// Create a copy to avoid modifying the original
	result := make([]int, len(arr))
	copy(result, arr)

	// First sort by value (ascending)
	sort.Ints(result)

	// Pre-compute count of 1s for each element
	countOnes := func(n int) int {
		count := 0
		for n != 0 {
			count += n & 1
			n >>= 1
		}
		return count
	}

	ones := make([]int, len(result))
	for i, v := range result {
		ones[i] = countOnes(v)
	}

	// Then sort by count of 1s in binary (ascending)
	// Using sort.SliceStable to preserve value order for equal 1-count elements
	sort.SliceStable(result, func(i, j int) bool {
		if ones[i] != ones[j] {
			return ones[i] < ones[j]
		}
		// If same count, sort by value (ascending)
		return result[i] < result[j]
	})

	return result
}

func main() {
	// Test cases
	test1 := []int{1, 5, 2, 3, 4}
	test2 := []int{-2, -3, -4, -5, -6}
	test3 := []int{1, 0, 2, 3, 4}

	println(SortArray(test1))
	println(SortArray(test2))
	println(SortArray(test3))
}
