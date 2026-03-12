package main

import (
	"fmt"
	"sort"
	"strconv"
)

func OrderByPoints(nums []int) []int {
	// Create a copy to avoid modifying original
	result := make([]int, len(nums))
	copy(result, nums)

	// Use sort.SliceStable to maintain original order for ties (stable sort)
	sort.SliceStable(result, func(i, j int) bool {
		return digitSum(result[i]) < digitSum(result[j])
	})

	return result
}

func digitSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}

	// Convert to string to get digits
	strN := strconv.Itoa(n)
	sum := 0
	for i, c := range strN {
		digit := int(c - '0')
		if i == 0 {
			digit *= neg
		}
		sum += digit
	}

	return sum
}

func main() {
	// Test the function
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12}))
	fmt.Println(OrderByPoints([]int{}))
}