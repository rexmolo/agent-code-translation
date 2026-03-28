package main

import (
    "fmt"
    "strings"
)

func CountUpper(s string) int {
    count := 0
    for i := 0; i < len(s); i += 2 {
        char := string(s[i])
        if strings.Contains("AEIOU", char) {
            count++
        }
    }
    return count
}

func main() {
    var s string
    fmt.Scan(&s)
    fmt.Println(CountUpper(s))
}