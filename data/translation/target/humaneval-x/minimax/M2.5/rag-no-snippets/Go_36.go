package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzz(n int) int {
	// Return the number of times the digit 7 appears in integers less than n which are divisible by 11 or 13.
	ns := []string{}
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			ns = append(ns, strconv.Itoa(i))
		}
	}
	s := strings.Join(ns, "")
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
