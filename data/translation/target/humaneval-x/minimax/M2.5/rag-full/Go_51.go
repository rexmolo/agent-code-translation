package main

import "strings"

func RemoveVowels(text string) string {
    const vowels = "aeiouAEIOU"
    var result []rune
    for _, char := range text {
        if strings.IndexByte(vowels, byte(char)) == -1 {
            result = append(result, char)
        }
    }
    return string(result)
}