package main

import "strings"

func DecodeShift(s string) string {
    result := make([]rune, len(s))
    for i, ch := range s {
        result[i] = rune(((int(ch) - 5 - int('a')) % 26) + int('a'))
    }
    return string(result)
}

func main() {
    // Example usage
    encoded := "fghjklmnopqrstuvwxyzabcde"
    decoded := DecodeShift(encoded)
    println(decoded) // Output: abcdefghijklmnopqrstuvwxyz
}
