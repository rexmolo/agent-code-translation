package main

import (
	"strconv"
)

func Specialfilter(nums []int) int {
	count := 0
	oddDigits := map[rune]bool{
		'1': true,
		'3': true,
		'5': true,
		'7': true,
		'9': true,
	}

	for _, num := range nums {
		if num > 10 {
			s := strconv.Itoa(num)
			firstDigit := rune(s[0])
			lastDigit := rune(s[len(s)-1])

			if oddDigits[firstDigit] && oddDigits[lastDigit] {
				count++
			}
		}
	}
	return count
}

func main() {
	// Test cases
	nums1 := []int{15, -73, 14, -15}
	nums2 := []int{33, -2, -3, 45, 21, 109}
	println(Specialfilter(nums1)) // Expected: 1
	println(Specialfilter(nums2)) // Expected: 2
}
