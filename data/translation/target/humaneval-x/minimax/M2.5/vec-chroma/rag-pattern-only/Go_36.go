package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) int {
	var ns []string
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			ns = append(ns, strconv.Itoa(i))
		}
	}
	s := ""
	for _, v := range ns {
		s += v
	}
	ans := 0
	for _, c := range s {
		if c == '7' {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(FizzBuzz(50)) // 0
	fmt.Println(FizzBuzz(78)) // 2
	fmt.Println(FizzBuzz(79)) // 3
}
