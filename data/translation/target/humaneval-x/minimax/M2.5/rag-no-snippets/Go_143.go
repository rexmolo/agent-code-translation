package main

import (
	"fmt"
	"strings"
)

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
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
	var newWords []string

	for _, word := range words {
		if isPrime(len(word)) {
			newWords = append(newWords, word)
		}
	}

	return strings.Join(newWords, " ")
}

func main() {
	fmt.Println(WordsInSentence("This is a test"))
	fmt.Println(WordsInSentence("lets go for swimming"))
}