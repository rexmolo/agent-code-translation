package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) int {
	// Collect all numbers less than n that are divisible by 11 or 13
	var ns []int
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			ns = append(ns, i)
		}
	}

	// Convert all numbers to strings and concatenate
	var s string
	for _, num := range ns {
		s += strconv.Itoa(num)
	}

	// Count occurrences of digit '7'
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