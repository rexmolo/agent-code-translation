package main

import (
	"fmt"
)

func Encode(message string) string {
	vowelsReplace := map[byte]byte{
		'a': 'c', 'e': 'g', 'i': 'k', 'o': 'q', 'u': 'w',
		'A': 'C', 'E': 'G', 'I': 'K', 'O': 'Q', 'U': 'W',
	}

	result := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		char := message[i]

		// Swap case
		if char >= 'a' && char <= 'z' {
			char -= 'a' - 'A'
		} else if char >= 'A' && char <= 'Z' {
			char += 'a' - 'A'
		}

		// Replace vowel if present in map
		if replacement, ok := vowelsReplace[char]; ok {
			char = replacement
		}

		result[i] = char
	}

	return string(result)
}

func main() {
	fmt.Println(Encode("test"))
	fmt.Println(Encode("This is a message"))
}