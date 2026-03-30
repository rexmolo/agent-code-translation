package main

import (
	"fmt"
	"sort"
)

func digitSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum * neg
}

func OrderByPoints(nums []int) []int {
	// Make a copy to avoid modifying the original slice
	result := make([]int, len(nums))
	copy(result, nums)

	// Use sort.SliceStable to maintain original order for equal elements
	// This matches Python's sorted() which is stable by default
	sort.SliceStable(result, func(i, j int) bool {
		return digitSum(result[i]) < digitSum(result[j])
	})

	return result
}

func main() {
	// Test examples
	fmt.Println(OrderByPoints([]int{}))                              // []
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12}))           // [-1 -11 1 -12 11]
}
