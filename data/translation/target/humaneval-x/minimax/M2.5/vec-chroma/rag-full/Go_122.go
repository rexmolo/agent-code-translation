package main

import (
	"fmt"
	"strconv"
)

func AddElements(arr []int, k int) int {
	sum := 0
	for i := 0; i < k && i < len(arr); i++ {
		if len(strconv.Itoa(arr[i])) <= 2 {
			sum += arr[i]
		}
	}
	return sum
}

func main() {
	arr := []int{111, 21, 3, 4000, 5, 6, 7, 8, 9}
	result := AddElements(arr, 4)
	fmt.Println(result)
}