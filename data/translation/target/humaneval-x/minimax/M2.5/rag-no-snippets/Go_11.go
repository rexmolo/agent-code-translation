package main

import (
    "strings"
)

func StringXor(a string, b string) string {
    xor := func(i, j rune) string {
        if i == j {
            return "0"
        }
        return "1"
    }

    result := make([]string, 0, min(len(a), len(b)))

    for i := 0; i < len(a) && i < len(b); i++ {
        result = append(result, xor(rune(a[i]), rune(b[i])))
    }

    return strings.Join(result, "")
}