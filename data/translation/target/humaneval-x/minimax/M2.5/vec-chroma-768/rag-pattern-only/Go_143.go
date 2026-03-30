package main

import (
    "fmt"
    "strings"
)

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func WordsInSentence(sentence string) string {
    words := strings.Split(sentence, " ")
    var result []string
    for _, word := range words {
        if isPrime(len(word)) {
            result = append(result, word)
        }
    }
    return strings.Join(result, " ")
}

func main() {
    test1 := "This is a test"
    fmt.Printf("Input: %q\nOutput: %q\n\n", test1, WordsInSentence(test1))

    test2 := "lets go for swimming"
    fmt.Printf("Input: %q\nOutput: %q\n\n", test2, WordsInSentence(test2))
}
