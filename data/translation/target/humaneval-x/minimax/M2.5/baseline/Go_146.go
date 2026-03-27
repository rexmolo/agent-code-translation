package main

import (
	"strconv"
	"strings"
)

func Specialfilter(nums []int) int {
	count := 0
	oddDigits := "13579"

	for _, num := range nums {
		if num > 10 {
			str := strconv.Itoa(num)
			firstDigit := string(str[0])
			lastDigit := string(str[len(str)-1])

			if strings.Contains(oddDigits, firstDigit) && strings.Contains(oddDigits, lastDigit) {
				count++
			}
		}
	}

	return count
}

func main() {
	// Test cases
	println(Specialfilter([]int{15, -73, 14, -15}))    // Expected: 1
	println(Specialfilter([]int{33, -2, -3, 45, 21, 109})) // Expected: 2
}