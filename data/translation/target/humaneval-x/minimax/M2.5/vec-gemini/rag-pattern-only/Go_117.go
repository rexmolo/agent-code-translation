package main

import (
    "fmt"
    "strings"
)

// SelectWords returns all words from string s that contain exactly n consonants,
// in the order they appear in the string s.
func SelectWords(s string, n int) []string {
    words := strings.Fields(s)
    
    vowels := map[rune]bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
        'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
    }
    
    var result []string
    for _, word := range words {
        nConsonants := 0
        for _, ch := range word {
            if !vowels[ch] {
                nConsonants++
            }
        }
        if nConsonants == n {
            result = append(result, word)
        }
    }
    
    return result
}

func main() {
    // Test cases
    fmt.Println(SelectWords("Mary had a little lamb", 4))   // [little]
    fmt.Println(SelectWords("Mary had a little lamb", 3))   // [Mary lamb]
    fmt.Println(SelectWords("simple white space", 2))        // []
    fmt.Println(SelectWords("Hello world", 4))              // [world]
    fmt.Println(SelectWords("Uncle sam", 3))                // [Uncle]
    fmt.Println(SelectWords("", 3))                          // []
}