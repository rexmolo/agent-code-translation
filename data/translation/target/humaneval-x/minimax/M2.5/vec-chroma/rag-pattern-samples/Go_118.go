package main

import "fmt"

func GetClosestVowel(word string) string {
	if len(word) < 3 {
		return ""
	}

	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	for i := len(word) - 2; i > 0; i-- {
		ch := rune(word[i])
		nextCh := rune(word[i+1])
		prevCh := rune(word[i-1])

		if vowels[ch] && !vowels[nextCh] && !vowels[prevCh] {
			return string(ch)
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
