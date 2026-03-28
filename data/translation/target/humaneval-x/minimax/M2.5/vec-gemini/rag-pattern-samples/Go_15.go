package main

import (
	"fmt"
	"strings"
)

func StringSequence(n int) string {
	parts := make([]string, n+1)
	for i := 0; i <= n; i++ {
		parts[i] = fmt.Sprintf("%d", i)
	}
	return strings.Join(parts, " ")
}