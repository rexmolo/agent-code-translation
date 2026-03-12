package main

import (
	"fmt"
	"sort"
)

func UniqueDigits(x []int) []int {
	var oddDigitElements []int

	for _, i := range x {
		num := i
		isAllOdd := true

		if num == 0 {
			// 0 is even, so skip
			isAllOdd = false
		} else {
			for num > 0 {
				digit := num % 10
				if digit%2 == 0 {
					isAllOdd = false
					break
				}
				num /= 10
			}
		}

		if isAllOdd {
			oddDigitElements = append(oddDigitElements, i)
		}
	}

	// Sort in increasing order
	sort.Slice(oddDigitElements, func(i, j int) bool {
		return oddDigitElements[i] < oddDigitElements[j]
	})

	return oddDigitElements
}

func main() {
	// Test cases
	result1 := UniqueDigits([]int{15, 33, 1422, 1})
	fmt.Println(result1) // [1 15 33]

	result2 := UniqueDigits([]int{152, 323, 1422, 10})
	fmt.Println(result2) // []
}
