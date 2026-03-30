package main

import (
    "strings"
)

func encodeShift(s string) string {
    var builder strings.Builder
    for _, ch := range s {
        encoded := ((int(ch-'a') + 5) % 26) + int('a')
        builder.WriteByte(byte(encoded))
    }
    return builder.String()
}

func decodeShift(s string) string {
    var builder strings.Builder
    for _, ch := range s {
        // Handle negative modulo in Go: ((x % 26) + 26) % 26
        decoded := (((int(ch-'a') - 5) % 26) + 26) % 26 + int('a')
        builder.WriteByte(byte(decoded))
    }
    return builder.String()
}

// For standalone testing
func main() {
    original := "hello"
    encoded := encodeShift(original)
    decoded := decodeShift(encoded)
    println("Original:", original)
    println("Encoded:", encoded)
    println("Decoded:", decoded)
}