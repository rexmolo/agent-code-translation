package main

import (
	"fmt"
	"strconv"
)

func Specialfilter(nums []int) int {
	count := 0
	for _, num := range nums {
		if num > 10 {
			oddDigits := map[int]bool{1: true, 3: true, 5: true, 7: true, 9: true}
			s := strconv.Itoa(num)
			first := int(s[0] - '0')
			last := int(s[len(s)-1] - '0')
			if oddDigits[first] && oddDigits[last] {
				count++
			}
		}
	}
	return count
}

func main() {
	// Test examples
	fmt.Println(Specialfilter([]int{15, -73, 14, -15}))       // 1
	fmt.Println(Specialfilter([]int{33, -2, -3, 45, 21, 109})) // 2
}