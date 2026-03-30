package main

import "fmt"

func GetClosestVowel(word string) string {
	if len(word) < 3 {
		return ""
	}

	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	for i := len(word) - 2; i > 0; i-- {
		if vowels[word[i]] {
			if !vowels[word[i+1]] && !vowels[word[i-1]] {
				return string(word[i])
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(GetClosestVowel("yogurt"))
	fmt.Println(GetClosestVowel("FULL"))
	fmt.Println(GetClosestVowel("quick"))
	fmt.Println(GetClosestVowel("ab"))
}
