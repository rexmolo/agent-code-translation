package main

import "fmt"

func DecodeShift(s string) string {
	// Create a rune slice to store the result
	result := make([]rune, len(s))

	for i, ch := range s {
		// Convert to lowercase if uppercase
		lowerCh := ch
		if ch >= 'A' && ch <= 'Z' {
			lowerCh = ch + ('a' - 'A')
		}

		// Apply the decoding: ((ord(ch) - 5 - ord('a')) % 26) + ord('a')
		// ((val % 26) + 26) % 26 ensures non-negative result for negative values
		shifted := ((int(lowerCh) - 5 - int('a')) % 26 + 26) % 26
		result[i] = rune(shifted + int('a'))
	}

	return string(result)
}

func main() {
	// Test the function
	encoded := "ujqhtwfz" // This is "example" shifted by 5
	decoded := DecodeShift(encoded)
	fmt.Println("Decoded:", decoded)
}
