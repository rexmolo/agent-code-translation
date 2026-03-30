package main

import "fmt"

func Specialfilter(nums []int) int {
	count := 0
	for _, num := range nums {
		if num > 10 {
			str := fmt.Sprintf("%d", num)
			// For negative numbers, first digit is at index 1 (after the minus sign)
			// For positive numbers, first digit is at index 0
			start := 0
			if str[0] == '-' {
				start = 1
			}
			first := int(str[start] - '0')
			last := int(str[len(str)-1] - '0')
			if first%2 == 1 && last%2 == 1 {
				count++
			}
		}
	}
	return count
}

func main() {
	// Test examples
	fmt.Println(Specialfilter([]int{15, -73, 14, -15})) // => 1
	fmt.Println(Specialfilter([]int{33, -2, -3, 45, 21, 109})) // => 2
}