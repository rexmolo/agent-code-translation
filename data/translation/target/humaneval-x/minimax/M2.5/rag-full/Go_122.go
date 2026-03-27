package main

import (
	"strconv"
)

func AddElements(arr []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		elem := arr[i]
		s := strconv.Itoa(elem)
		if len(s) <= 2 {
			sum += elem
		}
	}
	return sum
}
