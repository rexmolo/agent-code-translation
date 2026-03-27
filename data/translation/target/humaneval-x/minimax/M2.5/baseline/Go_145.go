package main

import (
	"fmt"
	"sort"
	"strconv"
)

func OrderByPoints(nums []int) []int {
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		s := strconv.Itoa(n)
		sum := 0
		for i, c := range s {
			d, _ := strconv.Atoi(string(c))
			if i == 0 {
				d *= neg
			}
			sum += d
		}
		return sum
	}

	result := make([]int, len(nums))
	copy(result, nums)

	sort.SliceStable(result, func(i, j int) bool {
		return digitsSum(result[i]) < digitsSum(result[j])
	})

	return result
}

func main() {
	fmt.Println(OrderByPoints([]int{}))
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12}))
}