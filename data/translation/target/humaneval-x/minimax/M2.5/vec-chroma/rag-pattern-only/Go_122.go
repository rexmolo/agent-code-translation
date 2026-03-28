package main

import (
	"strconv"
)

func AddElements(arr []int, k int) int {
	sum := 0
	// Iterate through first k elements (or fewer if k exceeds array length)
	for i := 0; i < k && i < len(arr); i++ {
		// Convert to string to check digit count
		if len(strconv.Itoa(arr[i])) <= 2 {
			sum += arr[i]
		}
	}
	return sum
}

func main() {
	// Example usage
	arr := []int{111, 21, 3, 4000, 5, 6, 7, 8, 9}
	k := 4
	result := AddElements(arr, k)
	print(result) // Output: 24
}
