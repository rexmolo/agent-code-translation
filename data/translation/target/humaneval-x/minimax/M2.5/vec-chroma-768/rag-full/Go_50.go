package main

import "fmt"

func DecodeShift(s string) string {
	var result []byte
	for _, ch := range s {
		encoded := (int(ch) - 5 - int('a') + 26) % 26 + int('a')
		result = append(result, byte(encoded))
	}
	return string(result)
}

func EncodeShift(s string) string {
	var result []byte
	for _, ch := range s {
		encoded := (int(ch) - int('a') + 5) % 26 + int('a')
		result = append(result, byte(encoded))
	}
	return string(result)
}

func main() {
	// Example usage
	original := "hello"
	encoded := EncodeShift(original)
	decoded := DecodeShift(encoded)
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Encoded: %s\n", encoded)
	fmt.Printf("Decoded: %s\n", decoded)
}
