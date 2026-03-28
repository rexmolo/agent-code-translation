package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func WordsInSentence(sentence string) string {
	words := strings.Split(sentence, " ")
	result := []string{}

	for _, word := range words {
		length := len(word)
		isNotPrime := 0

		if length == 1 {
			isNotPrime = 1
		}

		for i := 2; i < length; i++ {
			if length%i == 0 {
				isNotPrime = 1
				break
			}
		}

		if isNotPrime == 0 || length == 2 {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}

func main() {
	// Read input from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter sentence: ")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	result := WordsInSentence(sentence)
	fmt.Println(result)
}
