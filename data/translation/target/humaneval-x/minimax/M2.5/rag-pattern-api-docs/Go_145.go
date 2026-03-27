package main

import (
	"fmt"
	"sort"
	"strconv"
)

func OrderByPoints(nums []int) []int {
	// Create a copy to avoid modifying the original slice
	result := make([]int, len(nums))
	copy(result, nums)

	// Sort using custom key (digit sum)
	// sort.SliceStable maintains original order for equal keys
	sort.SliceStable(result, func(i, j int) bool {
		return digitSum(result[i]) < digitSum(result[j])
	})

	return result
}

// digitSum calculates the sum of digits of a number
// Negative numbers have their first digit multiplied by -1
func digitSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}

	// Convert number to string to get digits
	s := strconv.Itoa(n)

	sum := 0
	for i, c := range s {
		digit := int(c - '0')
		if i == 0 {
			digit *= neg // Multiply first digit by -1 for negative numbers
		}
		sum += digit
	}

	return sum
}

func main() {
	// Test cases
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12})) // Expected: [-1, -11, 1, -12, 11]
	fmt.Println(OrderByPoints([]int{}))                     // Expected: []
}
