package main

import (
	"fmt"
	"strconv"
)

func CountNums(arr []int) int {
	// Helper function to calculate digit sum with signed first digit
	digitsSum := func(n int) int {
		neg := 1
		if n < 0 {
			n = -n
			neg = -1
		}

		sum := 0
		strN := strconv.Itoa(n)
		for i, c := range strN {
			digit := int(c - '0')
			if i == 0 {
				digit = digit * neg
			}
			sum += digit
		}
		return sum
	}

	count := 0
	for _, n := range arr {
		if digitsSum(n) > 0 {
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(CountNums([]int{}))           // 0
	fmt.Println(CountNums([]int{-1, 11, -11})) // 1
	fmt.Println(CountNums([]int{1, 1, 2}))    // 3
}
