package main

import (
    "strings"
)

func RemoveVowels(text string) string {
    vowels := "aeiou"
    
    var result []rune
    for _, r := range text {
        if strings.Index(vowels, strings.ToLower(string(r))) == -1 {
            result = append(result, r)
        }
    }
    return string(result)
}
