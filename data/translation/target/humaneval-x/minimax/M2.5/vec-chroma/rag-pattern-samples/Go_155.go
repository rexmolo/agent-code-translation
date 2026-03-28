package main

import (
	"fmt"
	"strconv"
)

func EvenOddCount(num int) [2]int {
	evenCount := 0
	oddCount := 0
	n := num
	if n < 0 {
		n = -n
	}
	str := strconv.Itoa(n)
	for _, c := range str {
		digit := int(c - '0')
		if digit%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return [2]int{evenCount, oddCount}
}

func main() {
	fmt.Println(EvenOddCount(-12)) // [1 1]
	fmt.Println(EvenOddCount(123))  // [1 2]
}
