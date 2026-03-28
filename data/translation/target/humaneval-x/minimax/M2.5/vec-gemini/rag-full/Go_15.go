package main

import (
	"strconv"
	"strings"
)

func StringSequence(n int) string {
	strs := make([]string, n+1)
	for i := 0; i <= n; i++ {
		strs[i] = strconv.Itoa(i)
	}
	return strings.Join(strs, " ")
}