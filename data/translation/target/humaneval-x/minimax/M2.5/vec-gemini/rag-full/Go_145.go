package main

import (
	"fmt"
	"sort"
)

func OrderByPoints(nums []int) []int {
	// Helper to calculate digit sum
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		sum := 0
		if n == 0 {
			sum = 0
		} else {
			for n > 0 {
				sum += n % 10
				n /= 10
			}
		}
		return sum * neg
	}

	// Create a copy to avoid modifying original
	result := make([]int, len(nums))
	copy(result, nums)

	// Use sort.SliceStable to maintain original order for equal digit sums
	sort.SliceStable(result, func(i, j int) bool {
		return digitsSum(result[i]) < digitsSum(result[j])
	})

	return result
}

func main() {
	// Test cases
	test1 := []int{1, 11, -1, -11, -12}
	fmt.Println(OrderByPoints(test1)) // Should print: [-1 -11 1 -12 11]

	test2 := []int{}
	fmt.Println(OrderByPoints(test2)) // Should print: []
}