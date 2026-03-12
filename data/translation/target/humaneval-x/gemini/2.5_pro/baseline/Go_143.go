package main

import (
	"fmt"
	"strings"
)

// isPrime checks if a number n is a prime number.
// It returns true if n is prime, and false otherwise.
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	// A number is prime if it's not divisible by any number from 2 up to its square root.
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// WordsInSentence filters a sentence to include only words whose lengths are prime numbers.
// The order of the words is preserved.
func WordsInSentence(sentence string) string {
	// strings.Fields is the Go equivalent of Python's sentence.split(),
	// correctly handling various whitespace separators.
	words := strings.Fields(sentence)
	
	var primeWords []string

	for _, word := range words {
		if isPrime(len(word)) {
			primeWords = append(primeWords, word)
		}
	}

	// strings.Join is the Go equivalent of Python's " ".join().
	return strings.Join(primeWords, " ")
}

// main function to demonstrate the WordsInSentence function.
func main() {
	sentence1 := "This is a test"
	output1 := WordsInSentence(sentence1)
	fmt.Println(output1)

	sentence2 := "lets go for swimming"
	output2 := WordsInSentence(sentence2)
	fmt.Println(output2)
}
