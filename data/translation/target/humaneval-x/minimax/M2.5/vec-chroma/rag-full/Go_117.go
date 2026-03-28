package main

import (
    "strings"
)

// SelectWords returns all words from string s that contain exactly n consonants,
// in the order they appear in the string s.
func SelectWords(s string, n int) []string {
    vowels := map[string]bool{
        "a": true,
        "e": true,
        "i": true,
        "o": true,
        "u": true,
    }

    var result []string
    words := strings.Fields(s)

    for _, word := range words {
        nConsonants := 0
        for i := 0; i < len(word); i++ {
            char := strings.ToLower(string(word[i]))
            if !vowels[char] {
                nConsonants++
            }
        }
        if nConsonants == n {
            result = append(result, word)
        }
    }

    return result
}