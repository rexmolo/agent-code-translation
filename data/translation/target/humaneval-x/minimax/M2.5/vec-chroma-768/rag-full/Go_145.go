package main

import (
	"sort"
	"strconv"
)

func digitsSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}
	s := strconv.Itoa(n)
	sum := 0
	for i, c := range s {
		digit := int(c - '0')
		if i == 0 {
			digit *= neg
		}
		sum += digit
	}
	return sum
}

func OrderByPoints(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		sumI := digitsSum(nums[i])
		sumJ := digitsSum(nums[j])
		if sumI == sumJ {
			return i < j // tiebreaker: preserve original order (stable sort behavior)
		}
		return sumI < sumJ
	})
	return nums
}

func main() {
	// Test cases
	// result := OrderByPoints([]int{1, 11, -1, -11, -12})
	// fmt.Println(result)
}