package main

import (
	"fmt"
	"strconv"
)

func Specialfilter(nums []int) int {
	count := 0
	oddDigits := map[int]bool{1: true, 3: true, 5: true, 7: true, 9: true}
	
	for _, num := range nums {
		if num > 10 {
			numStr := strconv.Itoa(num)
			firstDigit := int(numStr[0] - '0')
			lastDigit := int(numStr[len(numStr)-1] - '0')
			
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
	
	fmt.Println(Specialfilter(nums1)) // Output: 1
	fmt.Println(Specialfilter(nums2)) // Output: 2
}
