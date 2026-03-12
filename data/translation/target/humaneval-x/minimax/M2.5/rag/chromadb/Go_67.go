package main

import (
	"strconv"
	"strings"
)

func FruitDistribution(s string, n int) int {
	parts := strings.Split(s, " ")
	sum := 0
	for _, part := range parts {
		if val, err := strconv.Atoi(part); err == nil {
			sum += val
		}
	}
	return n - sum
}