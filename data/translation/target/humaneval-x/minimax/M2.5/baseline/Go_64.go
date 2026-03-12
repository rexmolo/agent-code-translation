package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func VowelsCount(s string) int {
	vowels := "aeiouAEIOU"
	nVowels := 0

	for _, c := range s {
		if strings.Contains(vowels, string(c)) {
			nVowels++
		}
	}

	if len(s) > 0 {
		lastChar := s[len(s)-1]
		if lastChar == 'y' || lastChar == 'Y' {
			nVowels++
		}
	}

	return nVowels
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	result := VowelsCount(text)
	fmt.Println(result)
}
