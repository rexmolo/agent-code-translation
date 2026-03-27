package main

import (
	"strings"
)

func WordsInSentence(sentence string) string {
	words := strings.Split(sentence, " ")
	result := []string{}
	for _, word := range words {
		if isPrime(len(word)) {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
