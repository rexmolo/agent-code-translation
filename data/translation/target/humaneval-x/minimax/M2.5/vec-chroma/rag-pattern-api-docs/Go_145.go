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
			n = -1 * n
			neg = -1
		}
		// Convert number to string and get digit characters
		s := strconv.Itoa(n)
		chars := strings.Split(s, "")
		// Convert each character back to integer
		sum := 0
		for i, ch := range chars {
			d, _ := strconv.Atoi(ch)
			if i == 0 {
				d = d * neg
			}
			sum += d
		}
		return sum
	}

	// Sort with custom key using sort.Slice
	sort.Slice(nums, func(i, j int) bool {
		return digitsSum(nums[i]) < digitsSum(nums[j])
	})

	return nums
}

func main() {
	// Example usage
	result := OrderByPoints([]int{1, 11, -1, -11, -12})
	fmt.Println(result)

	result = OrderByPoints([]int{})
	fmt.Println(result)
}