package main

import (
	"strconv"
	"strings"
	"unicode"
)

func FruitDistribution(s string, n int) int {
	var total int
	parts := strings.Split(s, " ")
	for _, part := range parts {
		// Check if all characters in part are digits
		isDigit := true
		for _, r := range part {
			if !unicode.IsDigit(r) {
				isDigit = false
				break
			}
		}
		if isDigit && part != "" {
			num, _ := strconv.Atoi(part)
			total += num
		}
	}
	return n - total
}

func main() {
	// Test examples
	_ = FruitDistribution("5 apples and 6 oranges", 19)
	_ = FruitDistribution("0 apples and 1 oranges", 3)
	_ = FruitDistribution("2 apples and 3 oranges", 100)
	_ = FruitDistribution("100 apples and 1 oranges", 120)
}
