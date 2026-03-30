package main

import (
	"strconv"
	"strings"
)

func FizzBuzz(n int) int {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		if i%11 == 0 || i%13 == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
	}
	count := 0
	for _, c := range sb.String() {
		if c == '7' {
			count++
		}
	}
	return count
}