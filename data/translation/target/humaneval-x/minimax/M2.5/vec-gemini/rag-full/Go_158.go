package main

import (
    "fmt"
)

func FindMax(words []string) string {
    if len(words) == 0 {
        return ""
    }

    maxUnique := 0
    result := words[0]

    for _, word := range words {
        // Count unique characters using a map
        uniqueChars := make(map[rune]bool)
        for _, ch := range word {
            uniqueChars[ch] = true
        }

        uniqueCount := len(uniqueChars)

        // If this word has more unique characters, it's the new result
        // If equal unique count, check for lexicographically smaller
        if uniqueCount > maxUnique {
            maxUnique = uniqueCount
            result = word
        } else if uniqueCount == maxUnique && word < result {
            result = word
        }
    }

    return result
}

func main() {
    // Test cases
    fmt.Println(FindMax([]string{"name", "of", "string"}))    // "string"
    fmt.Println(FindMax([]string{"name", "enam", "game"}))   // "enam"
    fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"}))    // "bb"
}