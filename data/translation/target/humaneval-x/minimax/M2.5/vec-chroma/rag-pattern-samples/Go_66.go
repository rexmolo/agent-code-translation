package main

import "fmt"

func Digitsum(x string) int {
    if x == "" {
        return 0
    }
    
    sum := 0
    for _, c := range x {
        if c >= 'A' && c <= 'Z' {
            sum += int(c)
        }
    }
    return sum
}

func main() {
    // Test cases
    fmt.Println(Digitsum(""))       // 0
    fmt.Println(Digitsum("abAB"))   // 131
    fmt.Println(Digitsum("abcCd"))  // 67
    fmt.Println(Digitsum("helloE")) // 69
    fmt.Println(Digitsum("woArBld")) // 131
    fmt.Println(Digitsum("aAaaaXa")) // 153
}