package main

import (
	"strconv"
	"strings"
)

func StringSequence(n int) string {
	if n < 0 {
		return ""
	}
	parts := make([]string, n+1)
	for i := 0; i <= n; i++ {
		parts[i] = strconv.Itoa(i)
	}
	return strings.Join(parts, " ")
}