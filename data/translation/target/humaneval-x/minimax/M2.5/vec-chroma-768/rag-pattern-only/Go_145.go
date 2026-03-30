package main

import (
	"fmt"
	"sort"
)

// digitsSum calculates the sum of digits of a number.
// For negative numbers, only the first digit is treated as negative.
func digitsSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}

	sum := 0
	first := true
	for n > 0 {
		digit := n % 10
		if first {
			digit *= neg
			first = false
		}
		sum += digit
		n /= 10
	}

	return sum
}

func OrderByPoints(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// Create a copy to avoid modifying the original slice
	result := make([]int, len(nums))
	copy(result, nums)

	// Use stable sort to maintain original order for items with same digit sum
	sort.SliceStable(result, func(i, j int) bool {
		return digitsSum(result[i]) < digitsSum(result[j])
	})

	return result
}

func main() {
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12}))
	fmt.Println(OrderByPoints([]int{}))
}