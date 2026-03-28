package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzz(n int) int {
	var builder strings.Builder

	// Collect numbers divisible by 11 or 13
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			builder.WriteString(strconv.Itoa(i))
		}
	}

	// Count occurrences of digit '7'
	s := builder.String()
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