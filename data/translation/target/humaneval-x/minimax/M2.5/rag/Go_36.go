package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			s := strconv.Itoa(i)
			for _, c := range s {
				if c == '7' {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println(FizzBuzz(50))
	fmt.Println(FizzBuzz(78))
	fmt.Println(FizzBuzz(79))
}
