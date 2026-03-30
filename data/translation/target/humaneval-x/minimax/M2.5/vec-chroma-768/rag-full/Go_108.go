package main

import (
	"fmt"
	"strconv"
)

func CountNums(arr []int) int {
	count := 0
	for _, n := range arr {
		if digitsSum(n) > 0 {
			count++
		}
	}
	return count
}

func digitsSum(n int) int {
	neg := 1
	if n < 0 {
		n, neg = -n, -1
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
	// Test examples
	fmt.Println(CountNums([]int{}))             // 0
	fmt.Println(CountNums([]int{-1, 11, -11})) // 1
	fmt.Println(CountNums([]int{1, 1, 2}))     // 3
}
