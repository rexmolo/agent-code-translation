package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FruitDistribution(s string, n int) int {
	var nums []int
	for _, part := range strings.Split(s, " ") {
		if num, err := strconv.Atoi(part); err == nil {
			nums = append(nums, num)
		}
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return n - sum
}

func main() {
	fmt.Println(FruitDistribution("5 apples and 6 oranges", 19))
	fmt.Println(FruitDistribution("0 apples and 1 oranges", 3))
	fmt.Println(FruitDistribution("2 apples and 3 oranges", 100))
	fmt.Println(FruitDistribution("100 apples and 1 oranges", 120))
}
