package main

import "strings"

func VowelsCount(s string) int {
    vowels := "aeiouAEIOU"
    nVowels := 0

    for _, c := range s {
        if strings.ContainsRune(vowels, c) {
            nVowels++
        }
    }

    // Check if the last character is 'y' or 'Y'
    if len(s) > 0 {
        lastChar := s[len(s)-1]
        if lastChar == 'y' || lastChar == 'Y' {
            nVowels++
        }
    }

    return nVowels
}