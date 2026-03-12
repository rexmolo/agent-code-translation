package main

import "fmt"

func DecodeShift(s string) string {
	// Convert string to byte slice to manipulate each character
	bytes := []byte(s)
	for i, b := range bytes {
		// For each lowercase letter: subtract 'a' to get 0-25 range,
		// subtract 5 to decode, add 26 to handle negative modulo,
		// use % 26 to wrap around, then add 'a' back to get ASCII
		bytes[i] = ((b - 'a' - 5 + 26) % 26) + 'a'
	}
	return string(bytes)
}

func main() {
	// Example usage
	encoded := "fghjkl" // This is "abcde" shifted by 5
	decoded := DecodeShift(encoded)
	fmt.Println("Decoded:", decoded) // Should print: abcde
}
