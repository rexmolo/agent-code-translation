package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

func AntiShuffle(s string) string {
    words := strings.Split(s, " ")
    for i, word := range words {
        runes := []rune(word)
        sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
        words[i] = string(runes)
    }
    return strings.Join(words, " ")
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    fmt.Println(AntiShuffle(input))
}
