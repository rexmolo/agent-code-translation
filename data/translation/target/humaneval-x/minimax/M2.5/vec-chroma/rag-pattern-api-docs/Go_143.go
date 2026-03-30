package main

import (
    "strings"
)

func WordsInSentence(sentence string) string {
    words := strings.Split(sentence, " ")
    var newList []string
    
    for _, word := range words {
        wordLen := len(word)
        isNonPrime := false
        
        if wordLen == 1 {
            isNonPrime = true
        }
        for i := 2; i < wordLen; i++ {
            if wordLen%i == 0 {
                isNonPrime = true
                break
            }
        }
        if !isNonPrime {
            newList = append(newList, word)
        }
    }
    
    return strings.Join(newList, " ")
}

func main() {
    // Example usage
    println(WordsInSentence("This is a test"))    // Output: is
    println(WordsInSentence("lets go for swimming")) // Output: go for
}