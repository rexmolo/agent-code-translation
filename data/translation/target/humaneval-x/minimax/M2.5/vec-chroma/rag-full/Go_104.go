package main

import (
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var oddDigitElements []int
	for _, num := range x {
		if hasOnlyOddDigits(num) {
			oddDigitElements = append(oddDigitElements, num)
		}
	}
	sort.Ints(oddDigitElements)
	return oddDigitElements
}

func hasOnlyOddDigits(num int) bool {
	s := strconv.Itoa(num)
	for _, r := range s {
		digit := int(r - '0')
		if digit%2 == 0 {
			return false
		}
	}
	return true
}
