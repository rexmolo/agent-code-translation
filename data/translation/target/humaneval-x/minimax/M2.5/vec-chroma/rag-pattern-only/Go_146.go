package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Specialfilter(nums []int) int {
	count := 0
	oddDigits := "13579"

	for _, num := range nums {
		if num > 10 {
			numStr := strconv.Itoa(num)
			if strings.ContainsRune(oddDigits, rune(numStr[0])) &&
				strings.ContainsRune(oddDigits, rune(numStr[len(numStr)-1])) {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(Specialfilter([]int{15, -73, 14, -15}))
	fmt.Println(Specialfilter([]int{33, -2, -3, 45, 21, 109}))
}