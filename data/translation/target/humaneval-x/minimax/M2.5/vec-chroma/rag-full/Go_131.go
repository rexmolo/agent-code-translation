package main

import (
    "strconv"
)

func Digits(n int) int {
    product := 1
    oddCount := 0

    // Convert integer to string to iterate over each digit
    s := strconv.Itoa(n)

    for _, char := range s {
        digit := int(char - '0') // Convert rune to numeric digit
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

// Optional: main function to test the function
func main() {
    // Test cases
    println(Digits(1))  // Expected: 1
    println(Digits(4))  // Expected: 0
    println(Digits(235)) // Expected: 15
}