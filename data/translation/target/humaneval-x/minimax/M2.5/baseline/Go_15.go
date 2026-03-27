package main

import (
	"strconv"
	"strings"
)

func StringSequence(n int) string {
	var result []string
	for i := 0; i <= n; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return strings.Join(result, " ")
}