package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Test cases from docstring
	fmt.Println(CountNums([]int{}) == true)           // true
	fmt.Println(CountNums([]int{-1, 11, -11}) == 1)   // true
	fmt.Println(CountNums([]int{1, 1, 2}) == 3)       // true
}

func CountNums(arr []int) int {
	count := 0
	for _, n := range arr {
		if digitSum(n) > 0 {
			count++
		}
	}
	return count
}

// digitSum calculates the sum of digits with signed first digit
// If n is negative, the first digit becomes negative
func digitSum(n int) int {
	neg := 1
	if n < 0 {
		n = -n
		neg = -1
	}

	str := strconv.Itoa(n)
	sum := 0
	for i, c := range str {
		digit := int(c - '0')
		if i == 0 {
			digit *= neg
		}
		sum += digit
	}
	return sum
}
