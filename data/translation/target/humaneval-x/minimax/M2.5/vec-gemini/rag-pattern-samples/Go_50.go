package main

import "fmt"

func DecodeShift(s string) string {
	result := make([]rune, len(s))
	for i, ch := range s {
		// Convert to position in alphabet (0-25 where 'a' = 0)
		pos := int(ch) - int('a')
		// Shift back by 5 with modulo 26 (add 26 to handle negative values)
		pos = (pos - 5 + 26) % 26
		// Convert back to character
		result[i] = rune(pos + int('a'))
	}
	return string(result)
}

func main() {
	// Example usage
	encoded := "qfzqwv"
	decoded := DecodeShift(encoded)
	fmt.Println(decoded)
}
