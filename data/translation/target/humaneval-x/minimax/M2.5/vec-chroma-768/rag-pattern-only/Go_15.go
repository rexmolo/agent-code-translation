package main

import (
    "fmt"
    "strings"
)

func StringSequence(n int) string {
    strs := make([]string, 0, n+1)
    for i := 0; i <= n; i++ {
        strs = append(strs, fmt.Sprintf("%d", i))
    }
    return strings.Join(strs, " ")
}