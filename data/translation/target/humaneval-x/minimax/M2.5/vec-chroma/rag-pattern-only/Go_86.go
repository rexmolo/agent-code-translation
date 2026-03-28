package main

import (
    "fmt"
    "sort"
    "strings"
)

func AntiShuffle(s string) string {
    words := strings.Split(s, " ")
    
    for i, word := range words {
        runes := []rune(word)
        sort.Slice(runes, func(i, j int) bool {
            return runes[i] < runes[j]
        })
        words[i] = string(runes)
    }
    
    return strings.Join(words, " ")
}

func main() {
    fmt.Println(AntiShuffle("Hi"))
    fmt.Println(AntiShuffle("hello"))
    fmt.Println(AntiShuffle("Hello World!!!"))
}
