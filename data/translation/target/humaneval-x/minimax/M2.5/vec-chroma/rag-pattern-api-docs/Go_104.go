package main

import (
	"sort"
	"strconv"
)

func UniqueDigits(x []int) []int {
	var result []int
	for _, n := range x {
		str := strconv.Itoa(n)
		allOdd := true
		for _, r := range str {
			digit := int(r - '0')
			if digit%2 == 0 {
				allOdd = false
				break
			}
		}
		if allOdd {
			result = append(result, n)
		}
	}
	sort.Ints(result)
	return result
}