package main

import (
	"fmt"
	"math"
)

func Specialfilter(nums []int) int {
	count := 0
	oddDigits := map[int]bool{1: true, 3: true, 5: true, 7: true, 9: true}

	for _, num := range nums {
		if num > 10 {
			absNum := int(math.Abs(float64(num)))
			s := fmt.Sprintf("%d", absNum)
			firstDigit := int(s[0] - '0')
			lastDigit := int(s[len(s)-1] - '0')

			if oddDigits[firstDigit] && oddDigits[lastDigit] {
				count++
			}
		}
	}

	return count
}