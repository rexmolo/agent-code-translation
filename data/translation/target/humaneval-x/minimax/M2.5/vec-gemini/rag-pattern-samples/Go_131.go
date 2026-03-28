package main

import (
    "fmt"
    "strconv"
)

func Digits(n int) int {
    product := 1
    oddCount := 0

    strN := strconv.Itoa(n)
    for _, char := range strN {
        digit := int(char - '0')
        if digit%2 == 1 {
            product *= digit
            oddCount++
        }
    }

    if oddCount == 0 {
        return 0
    }
    return product
}

func main() {
    // Test cases
    fmt.Println(Digits(1))   // == 1
    fmt.Println(Digits(4))   // == 0
    fmt.Println(Digits(235)) // == 15
}
