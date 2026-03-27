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
    fmt.Println(Digitsum(""))
    fmt.Println(Digitsum("abAB"))
    fmt.Println(Digitsum("abcCd"))
    fmt.Println(Digitsum("helloE"))
    fmt.Println(Digitsum("woArBld"))
    fmt.Println(Digitsum("aAaaaXa"))
}
