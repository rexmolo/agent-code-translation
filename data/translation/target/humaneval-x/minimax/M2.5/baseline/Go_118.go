package main

import (
	"fmt"
	"strings"
)

func GetClosestVowel(word string) string {
	if len(word) < 3 {
		return ""
	}

	vowels := "aeiouAEIOU"

	for i := len(word) - 2; i > 0; i-- {
		char := word[i]
		if strings.Contains(vowels, string(char)) {
			left := word[i-1]
			right := word[i+1]
			if !strings.Contains(vowels, string(left)) && !strings.Contains(vowels, string(right)) {
				return string(word[i])
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(GetClosestVowel("yogurt")) // "u"
	fmt.Println(GetClosestVowel("FULL"))  // "U"
	fmt.Println(GetClosestVowel("quick")) // ""
	fmt.Println(GetClosestVowel("ab"))    // ""
}
