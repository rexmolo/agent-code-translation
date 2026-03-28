package main

import (
	"fmt"
	"strconv"
)

func Specialfilter(nums []int) int {
	count := 0

	for _, num := range nums {
		if num > 10 {
			numberAsString := strconv.Itoa(num)
			firstDigit := int(numberAsString[0] - '0')
			lastDigit := int(numberAsString[len(numberAsString)-1] - '0')

			// Check if both first and last digits are odd (1, 3, 5, 7, 9)
			firstIsOdd := firstDigit%2 == 1 && firstDigit != 5
			lastIsOdd := lastDigit%2 == 1 && lastDigit != 5

			// Alternatively, check against specific odd digits
			if (firstDigit == 1 || firstDigit == 3 || firstDigit == 5 || firstDigit == 7 || firstDigit == 9) &&
				(lastDigit == 1 || lastDigit == 3 || lastDigit == 5 || lastDigit == 7 || lastDigit == 9) {
				count++
			}
		}
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(Specialfilter([]int{15, -73, 14, -15}))             // Expected: 1
	fmt.Println(Specialfilter([]int{33, -2, -3, 45, 21, 109}))       // Expected: 2
}
