package main

import (
	"fmt"
	"sort"
)

func OrderByPoints(nums []int) []int {
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		sum := 0
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		// The Python logic multiplies the first digit by neg
		// Since we're summing from right to left (last digit first),
		// we need to handle this differently
		// Actually, let's use string conversion which is simpler
		return 0 // placeholder, will be recalculated below
	}

	// For proper translation, let's implement it properly
	// The Python logic: convert to string, treat each char as digit,
	// multiply first digit by neg, then sum

	type item struct {
		value    int
		digitsum int
	}

	items := make([]item, len(nums))
	for i, n := range nums {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}
		// Convert to string and sum digits
		s := fmt.Sprintf("%d", n)
		sum := 0
		for j, c := range s {
			d := int(c - '0')
			if j == 0 {
				d *= neg
			}
			sum += d
		}
		items[i] = item{value: nums[i], digitsum: sum}
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].digitsum != items[j].digitsum {
			return items[i].digitsum < items[j].digitsum
		}
		return i < j // tiebreaker: original index
	})

	result := make([]int, len(nums))
	for i, it := range items {
		result[i] = it.value
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(OrderByPoints([]int{1, 11, -1, -11, -12}))
	fmt.Println(OrderByPoints([]int{}))
}