package main

import (
    "fmt"
    "strings"
)

func SelectWords(s string, n int) []string {
    // Handle empty string case
    if s == "" {
        return []string{}
    }
    
    words := strings.Split(s, " ")
    result := []string{}
    
    for _, word := range words {
        n_consonants := 0
        for i := 0; i < len(word); i++ {
            ch := strings.ToLower(string(word[i]))
            // Check if character is not a vowel (consonant)
            if ch != "a" && ch != "e" && ch != "i" && ch != "o" && ch != "u" {
                n_consonants++
            }
        }
        if n_consonants == n {
            result = append(result, word)
        }
    }
    
    return result
}

func main() {
    // Test cases from the examples
    fmt.Println(SelectWords("Mary had a little lamb", 4)) // [little]
    fmt.Println(SelectWords("Mary had a little lamb", 3)) // [Mary lamb]
    fmt.Println(SelectWords("simple white space", 2))     // []
    fmt.Println(SelectWords("Hello world", 4))            // [world]
    fmt.Println(SelectWords("Uncle sam", 3))              // [Uncle]
    fmt.Println(SelectWords("", 3))                       // []
}
