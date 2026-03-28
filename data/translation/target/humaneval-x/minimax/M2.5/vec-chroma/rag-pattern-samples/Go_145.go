package main

import (
	"fmt"
	"sort"
)

func digitsSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}
	// Handle 0 case
	if n == 0 {
		return 0
	}
	sum := 0
	firstDigit := true
	for n > 0 {
		digit := n % 10
		if firstDigit {
			sum += digit * neg
			firstDigit = false
		} else {
			sum += digit
		}
		n /= 10
	}
	return sum
}

func OrderByPoints(nums []int) []int {
	// Create a copy to preserve indices for stable sorting
	result := make([]int, len(nums))
	copy(result, nums)

	sort.SliceStable(result, func(i, j int) bool {
		return digitsSum(result[i]) < digitsSum(result[j])
	})

	return result
}

func main() {
	// Test cases
	fmt.Println(OrderByPoints([]int{})) // []
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12})) // [-1, -11, 1, -12, 11]
}
