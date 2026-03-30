package main

import (
    "fmt"
    "strings"
    "unicode"
)

func CheckIfLastCharIsALetter(txt string) bool {
    parts := strings.Split(txt, " ")
    if len(parts) == 0 {
        return false
    }
    last := parts[len(parts)-1]
    if len(last) == 1 && unicode.IsLetter(rune(last[0])) {
        return true
    }
    return false
}

func main() {
    fmt.Println(CheckIfLastCharIsALetter("apple pie"))     // false
    fmt.Println(CheckIfLastCharIsALetter("apple pi e"))    // true
    fmt.Println(CheckIfLastCharIsALetter("apple pi e "))  // false
    fmt.Println(CheckIfLastCharIsALetter(""))             // false
}