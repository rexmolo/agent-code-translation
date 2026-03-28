package main

import (
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var result []int

	for _, num := range x {
		// Convert number to string to iterate over digits
		str := strconv.Itoa(num)
		allOdd := true
		for _, c := range str {
			digit := int(c - '0')
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