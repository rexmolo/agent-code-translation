package main

import (
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	oddDigitElements := []int{}

	for _, i := range x {
		if allOddDigits(i) {
			oddDigitElements = append(oddDigitElements, i)
		}
	}

	sort.Ints(oddDigitElements)
	return oddDigitElements
}

func allOddDigits(n int) bool {
	// Handle 0: it has an even digit (0), so it's not odd
	if n == 0 {
		return false
	}

	str := strconv.Itoa(n)
	for _, c := range str {
		digit := c - '0'
		if digit%2 == 0 {
			return false
		}
	}
	return true
}
