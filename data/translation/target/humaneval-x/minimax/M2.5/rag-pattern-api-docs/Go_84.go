package main

import (
    "fmt"
    "strconv"
)

func Solve(N int) string {
    // Convert N to string to iterate over each digit
    strN := strconv.Itoa(N)
    
    // Sum all digits
    sum := 0
    for _, c := range strN {
        digit := int(c - '0')
        sum += digit
    }
    
    // Convert sum to binary string (without "0b" prefix)
    return fmt.Sprintf("%b", sum)
}
