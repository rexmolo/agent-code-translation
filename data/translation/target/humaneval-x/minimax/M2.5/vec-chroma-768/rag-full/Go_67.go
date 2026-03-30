package main

import (
    "strings"
    "unicode"
)

func FruitDistribution(s string, n int) int {
    parts := strings.Split(s, " ")
    var sum int
    for _, part := range parts {
        if part == "" {
            continue
        }
        isDigit := true
        for _, r := range part {
            if !unicode.IsDigit(r) {
                isDigit = false
                break
            }
        }
        if isDigit {
            // Convert string to int using basic string parsing
            // since strconv would error on empty string which we already handled
            num := 0
            for _, r := range part {
                num = num*10 + int(r-'0')
            }
            sum += num
        }
    }
    return n - sum
}