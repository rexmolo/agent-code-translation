package main

import "fmt"

func DecodeShift(s string) string {
    result := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        // Shift back by 5 (subtract 5), wrapping around the alphabet
        result[i] = byte(((int(s[i]) - 5 - int('a')) % 26) + int('a'))
    }
    return string(result)
}

func main() {
    // Test decode
    encoded := "khoor" // "hello" encoded with shift +5
    decoded := DecodeShift(encoded)
    fmt.Println(decoded)
}
