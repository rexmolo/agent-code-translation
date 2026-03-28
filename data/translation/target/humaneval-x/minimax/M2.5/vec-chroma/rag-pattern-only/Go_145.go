package main

import (
	"fmt"
	"sort"
	"strconv"
)

func OrderByPoints(nums []int) []int {
	// Helper function to calculate sum of digits
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		// Convert number to string to get individual digits
		s := strconv.Itoa(n)
		// Calculate sum of digits
		sum := 0
		first := true
		for _, c := range s {
			d, _ := strconv.Atoi(string(c))
			if first {
				d *= neg
				first = false
			}
			sum += d
		}
		return sum
	}

	// Create items with index to track original positions for stable sorting
	type item struct {
		value int
		idx   int
	}
	items := make([]item, len(nums))
	for i, v := range nums {
		items[i] = item{value: v, idx: i}
	}

	// Use SliceStable to maintain original order for equal digit sums
	sort.SliceStable(items, func(i, j int) bool {
		return digitsSum(items[i].value) < digitsSum(items[j].value)
	})

	// Extract sorted values
	result := make([]int, len(items))
	for i, it := range items {
		result[i] = it.value
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12})) // Expected: [-1, -11, 1, -12, 11]
	fmt.Println(OrderByPoints([]int{}))                      // Expected: []
}