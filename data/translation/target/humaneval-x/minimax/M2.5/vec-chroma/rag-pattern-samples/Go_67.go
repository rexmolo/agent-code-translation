package main

import (
    "strconv"
    "strings"
)

func FruitDistribution(s string, n int) int {
    sum := 0
    parts := strings.Split(s, " ")
    for _, part := range parts {
        if num, err := strconv.Atoi(part); err == nil {
            sum += num
        }
    }
    return n - sum
}
