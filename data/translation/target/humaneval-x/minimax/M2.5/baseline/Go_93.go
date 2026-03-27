package main

import "fmt"

func Encode(message string) string {
    vowels := "aeiouAEIOU"
    result := make([]byte, len(message))
    
    for i, char := range message {
        // Check if character is a vowel
        if contains(vowels, char) {
            // Replace vowel with char + 2
            result[i] = byte(char + 2)
        } else {
            // Swap case
            if char >= 'a' && char <= 'z' {
                result[i] = byte(char - 'a' + 'A')
            } else if char >= 'A' && char <= 'Z' {
                result[i] = byte(char - 'A' + 'a')
            } else {
                result[i] = byte(char)
            }
        }
    }
    return string(result)
}

func contains(s string, char rune) bool {
    for _, c := range s {
        if c == char {
            return true
               }
    }
    return false
}

func main() {
    fmt.Println(Encode("test"))
    fmt.Println(Encode("This is a message"))
}
