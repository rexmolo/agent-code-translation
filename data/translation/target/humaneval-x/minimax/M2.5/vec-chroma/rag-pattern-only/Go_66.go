package main

import (
    "fmt"
    "unicode"
)

func Digitsum(x string) int {
    sum := 0
    for _, r := range x {
        if unicode.IsUpper(r) {
            sum += int(r)
        }
    }
    return sum
}

func main() {
    // Test the function
    fmt.Println(Digitsum(""))        // Expected: 0
    fmt.Println(Digitsum("abAB"))    // Expected: 131
    fmt.Println(Digitsum("abcCd"))   // Expected: 67
    fmt.Println(Digitsum("helloE"))  // Expected: 69
    fmt.Println(Digitsum("woArBld")) // Expected: 131
    fmt.Println(Digitsum("aAaaaXa")) // Expected: 153
}