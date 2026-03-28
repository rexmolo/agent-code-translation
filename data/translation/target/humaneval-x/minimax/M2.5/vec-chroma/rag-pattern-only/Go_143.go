package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPrime(n int) bool {
	if n <= 1 {
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter sentence: ")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	output := WordsInSentence(sentence)
	fmt.Println("Output:", output)
}
