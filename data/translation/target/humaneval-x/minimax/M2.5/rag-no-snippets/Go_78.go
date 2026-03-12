package main

import (
    "fmt"
    "strings"
)

func HexKey(num string) int {
    primes := "2357BD"
    total := 0
    for _, c := range num {
        if strings.Contains(primes, string(c)) {
            total++
        }
    }
    return total
}

func main() {
    fmt.Println(HexKey("AB"))           // 1
    fmt.Println(HexKey("1077E"))        // 2
    fmt.Println(HexKey("ABED1A33"))    // 4
    fmt.Println(HexKey("123456789ABCDEF0")) // 6
    fmt.Println(HexKey("2020"))        // 2
}