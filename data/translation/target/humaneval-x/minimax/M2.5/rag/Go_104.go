package main

import (
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var oddDigitElements []int

	for _, i := range x {
		// Convert number to string to check each digit
		s := strconv.Itoa(i)
		allOdd := true
		for _, c := range s {
			digit := int(c - '0')
			if digit%2 == 0 {
				allOdd = false
				break
			}
		}
		if allOdd {
			oddDigitElements = append(oddDigitElements, i)
		}
	}

	sort.Ints(oddDigitElements)
	return oddDigitElements
}
