package main

import (
	"fmt"
	"strconv"
)

func Specialfilter(nums []int) int {
	count := 0
	for _, num := range nums {
		if num > 10 {
			s := strconv.Itoa(num)
			first := s[0]
			last := s[len(s)-1]
			// Skip negative sign if present
			if first == '-' {
				first = s[1]
			}
			// Check if both first and last digits are odd (1, 3, 5, 7, 9)
			if (first-'0')%2 == 1 && (last-'0')%2 == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(Specialfilter([]int{15, -73, 14, -15}))          // Output: 1
	fmt.Println(Specialfilter([]int{33, -2, -3, 45, 21, 109}))   // Output: 2
}
