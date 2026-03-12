package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func OrderByPoints(nums []int) []int {
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		// Convert to string and get digits
		str := strconv.Itoa(n)
		var sum int
		for i, c := range str {
			d, _ := strconv.Atoi(string(c))
			if i == 0 {
				sum += d * neg // Multiply first digit by neg to handle negatives
			} else {
				sum += d
			}
		}
		return sum
	}

	// Create a copy to avoid modifying the original
	result := make([]int, len(nums))
	copy(result, nums)

	// Use SliceStable to preserve original order for equal digit sums
	sort.SliceStable(result, func(i, j int) bool {
		return digitsSum(result[i]) < digitsSum(result[j])
	})

	return result
}

func main() {
	// Test cases
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12})) // Expected: [-1, -11, 1, -12, 11]
	fmt.Println(OrderByPoints([]int{}))                     // Expected: []
}
