package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzz(n int) int {
	var ns []int
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			ns = append(ns, i)
		}
	}

	// Convert each integer to string and collect in a slice
	var strs []string
	for _, v := range ns {
		strs = append(strs, strconv.Itoa(v))
	}

	// Join all strings together
	s := strings.Join(strs, "")

	// Count occurrences of '7'
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
