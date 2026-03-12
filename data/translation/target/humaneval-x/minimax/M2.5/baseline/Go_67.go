package main

import (
	"strconv"
	"strings"
)

func FruitDistribution(s string, n int) int {
	sum := 0
	parts := strings.Split(s, " ")
	for _, part := range parts {
		if isDigitString(part) {
			num, _ := strconv.Atoi(part)
			sum += num
		}
	}
	return n - sum
}

func isDigitString(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
