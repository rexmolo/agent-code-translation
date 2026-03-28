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

	// Convert each integer to string and join them
	var strs []string
	for _, num := range ns {
		strs = append(strs, strconv.Itoa(num))
	}
	s := strings.Join(strs, "")

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
