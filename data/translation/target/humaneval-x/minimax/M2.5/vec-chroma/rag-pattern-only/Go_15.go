package main

import (
	"strconv"
	"strings"
)

func StringSequence(n int) string {
	var parts []string
	for i := 0; i <= n; i++ {
		parts = append(parts, strconv.Itoa(i))
	}
	return strings.Join(parts, " ")
}

func main() {
	// Test cases
	println(StringSequence(0))
	println(StringSequence(5))
}