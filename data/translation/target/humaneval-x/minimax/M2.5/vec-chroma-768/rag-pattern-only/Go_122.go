package main

import (
	"fmt"
	"math"
	"strconv"
)

func AddElements(arr []int, k int) int {
	sum := 0
	for i := 0; i < k && i < len(arr); i++ {
		elem := arr[i]
		// Convert absolute value to string to count digits
		// math.Abs works with float64, so we need to convert
		str := strconv.Itoa(int(math.Abs(float64(elem))))
		if len(str) <= 2 {
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
	fmt.Println(result) // Output: 24
}
