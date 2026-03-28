package main

import "fmt"

// DecodeShift takes as input string encoded with encode_shift function.
// Returns decoded string by shifting every character back by 5 in the alphabet.
func DecodeShift(s string) string {
	// Pre-allocate slice with capacity to avoid reallocations
	result := make([]byte, 0, len(s))
	
	for _, ch := range s {
		// Convert char to 0-25 range, subtract 5, wrap around using modulo
		// Add 26 before modulo to handle negative results in Go
		shifted := (int(ch) - 5 - int('a') + 26) % 26
		result = append(result, byte(shifted + int('a')))
	}
	
	return string(result)
}

func main() {
	// Example usage
	encoded := "mjqqt" // This is "hello" shifted by 5
	decoded := DecodeShift(encoded)
	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decoded)
}