package main

import "fmt"

func EncodeShift(s string) string {
    // Returns encoded string by shifting every character by 5 in the alphabet.
    result := make([]rune, 0, len(s))
    for _, ch := range s {
        shifted := (int(ch) - int('a') + 5) % 26
        result = append(result, rune(shifted+int('a')))
    }
    return string(result)
}

func DecodeShift(s string) string {
    // Takes as input string encoded with EncodeShift function. Returns decoded string.
    result := make([]rune, 0, len(s))
    for _, ch := range s {
        shifted := (int(ch) - 5 - int('a')) % 26
        if shifted < 0 {
            shifted += 26
        }
        result = append(result, rune(shifted+int('a')))
    }
    return string(result)
}

func main() {
    // Example usage
    original := "hello"
    encoded := EncodeShift(original)
    decoded := DecodeShift(encoded)
    fmt.Println("Original:", original)
    fmt.Println("Encoded:", encoded)
    fmt.Println("Decoded:", decoded)
}
