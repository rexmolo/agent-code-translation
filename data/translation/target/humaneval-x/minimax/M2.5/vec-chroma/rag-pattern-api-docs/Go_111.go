package main

import (
    "fmt"
    "strings"
)

func Histogram(test string) map[rune]int {
    if test == "" {
        return map[rune]int{}
    }

    parts := strings.Split(test, " ")

    // Count occurrences of each letter
    counts := make(map[string]int)
    for _, part := range parts {
        if part != "" {
            counts[part]++
        }
    }

    // Find the maximum count
    maxCount := 0
    for _, count := range counts {
        if count > maxCount {
            maxCount = count
        }
    }

    // Build result with letters that have the maximum count
    result := make(map[rune]int)
    for letter, count := range counts {
        if count == maxCount {
            // Convert string to rune (take first rune since it's a single character)
            runes := []rune(letter)
            if len(runes) > 0 {
                result[runes[0]] = count
            }
        }
    }

    return result
}

func main() {
    // Test cases
    fmt.Println(Histogram("a b c"))      // map[a:1 b:1 c:1]
    fmt.Println(Histogram("a b b a"))    // map[a:2 b:2]
    fmt.Println(Histogram("a b c a b")) // map[a:2 b:2]
    fmt.Println(Histogram("b b b b a")) // map[b:4]
    fmt.Println(Histogram(""))          // map[]
}