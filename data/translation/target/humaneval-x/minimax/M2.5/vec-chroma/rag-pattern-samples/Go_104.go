package main

import (
	"fmt"
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var result []int

	for _, num := range x {
		// Convert number to string to check each digit
		strNum := strconv.Itoa(num)
		allOdd := true

		for _, c := range strNum {
			digit, _ := strconv.Atoi(string(c))
			if digit%2 == 0 {
				allOdd = false
				break
			}
		}

		if allOdd {
			result = append(result, num)
		}
	}

	sort.Ints(result)
	return result
}

func main() {
	// Test case 1
	result1 := UniqueDigits([]int{15, 33, 1422, 1})
	fmt.Println(result1)

	// Test case 2
	result2 := UniqueDigits([]int{152, 323, 1422, 10})
	fmt.Println(result2)
}
