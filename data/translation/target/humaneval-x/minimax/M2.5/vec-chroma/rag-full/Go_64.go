package main

import (
    "fmt"
    "strings"
)

func VowelsCount(s string) int {
    vowels := "aeiouAEIOU"
    nVowels := 0

    for _, c := range s {
        if strings.Contains(vowels, string(c)) {
            nVowels++
        }
    }

    // Check if the last character is 'y' or 'Y'
    if len(s) > 0 && (s[len(s)-1] == 'y' || s[len(s)-1] == 'Y') {
        nVowels++
    }

    return nVowels
}

func main() {
    fmt.Println(VowelsCount("abcde"))
    fmt.Println(VowelsCount("ACEDY"))
}