package main

import (
    "fmt"
    "strconv"
)

func Solve(N int) string {
    // Convert N to string to iterate over each digit
    s := strconv.Itoa(N)
    
    // Sum all digits by iterating over each character
    sum := 0
    for _, c := range s {
        sum += int(c - '0')
    }
    
    // Convert sum to binary string using fmt.Sprintf
    return fmt.Sprintf("%b", sum)
}

func main() {
    // Example usage
    examples := []int{1000, 150, 147, 0}
    for _, n := range examples {
        fmt.Printf("N=%d -> %s\n", n, Solve(n))
    }
}