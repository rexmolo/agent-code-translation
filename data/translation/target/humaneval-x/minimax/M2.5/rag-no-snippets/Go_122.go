package main

import (
	"strconv"
)

func AddElements(arr []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		elem := arr[i]
		if len(strconv.Itoa(elem)) <= 2 {
			sum += elem
		}
	}
	return sum
}

func main() {
	// Example test
	arr := []int{111, 21, 3, 4000, 5, 6, 7, 8, 9}
	k := 4
	result := AddElements(arr, k)
	println(result) // Expected output: 24
}