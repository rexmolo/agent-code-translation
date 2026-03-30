package main

import (
	"strings"
)

func WordsInSentence(sentence string) string {
	words := strings.Split(sentence, " ")
	var newList []string
	for _, word := range words {
		length := len(word)
		if isPrime(length) {
			newList = append(newList, word)
		}
	}
	return strings.Join(newList, " ")
}

func isPrime(n int) bool {
	if n <= 1 {
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
