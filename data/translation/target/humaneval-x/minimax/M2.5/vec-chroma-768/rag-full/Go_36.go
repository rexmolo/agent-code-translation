package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzz(n int) int {
	var nums []int
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			nums = append(nums, i)
		}
	}

	var parts []string
	for _, num := range nums {
		parts = append(parts, strconv.Itoa(num))
	}
	s := strings.Join(parts, "")

	ans := 0
	for _, c := range s {
		if c == '7' {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(FizzBuzz(50))
	fmt.Println(FizzBuzz(78))
	fmt.Println(FizzBuzz(79))
}
