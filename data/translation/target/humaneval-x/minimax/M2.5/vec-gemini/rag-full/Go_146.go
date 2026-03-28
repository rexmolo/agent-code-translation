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
			numberAsString := strconv.Itoa(num)
			firstDigit := int(numberAsString[0] - '0')
			lastDigit := int(numberAsString[len(numberAsString)-1] - '0')

			if oddDigits[firstDigit] && oddDigits[lastDigit] {
				count++
			}
		}
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(Specialfilter([]int{15, -73, 14, -15}))   // Output: 1
	fmt.Println(Specialfilter([]int{33, -2, -3, 45, 21, 109})) // Output: 2
}
