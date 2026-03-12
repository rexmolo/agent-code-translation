package main

import (
	"fmt"
	"sort"
	"strconv"
)

func OrderByPoints(nums []int) []int {
	digitSums := make([]int, len(nums))
	for i, n := range nums {
		digitSums[i] = digitsSum(n)
	}

	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)

	sort.SliceStable(sortedNums, func(i, j int) bool {
		if digitSums[i] != digitSums[j] {
			return digitSums[i] < digitSums[j]
		}
		return false
	})

	return sortedNums
}

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

func main() {
	fmt.Println(OrderByPoints([]int{}))
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12}))
}