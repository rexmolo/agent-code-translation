package main

import (
	"fmt"
	"strings"
)

// isPrime checks if a number is prime
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
	var result []string
	
	for _, word := range words {
		if isPrime(len(word)) {
			result = append(result, word)
		}
	}
	
	return strings.Join(result, " ")
}

func main() {
	// Test examples
	fmt.Println(WordsInSentence("This is a test"))           // Output: "is"
	fmt.Println(WordsInSentence("lets go for swimming"))    // Output: "go for"
}
