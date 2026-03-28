package main

import (
    "strconv"
    "strings"
)

func FruitDistribution(s string, n int) int {
    parts := strings.Split(s, " ")
    sum := 0
    for _, part := range parts {
        // Try to convert each part to an integer
        // If successful, it's a number (apples or oranges count)
        if num, err := strconv.Atoi(part); err == nil {
            sum += num
        }
    }
    return n - sum
}
