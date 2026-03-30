package main

import (
    "strconv"
    "strings"
)

func FruitDistribution(s string, n int) int {
    parts := strings.Split(s, " ")
    total := 0
    
    for _, part := range parts {
        if num, err := strconv.Atoi(part); err == nil {
            total += num
        }
    }
    
    return n - total
}

func main() {
    // Test examples
    println(FruitDistribution("5 apples and 6 oranges", 19))
    println(FruitDistribution("0 apples and 1 oranges", 3))
    println(FruitDistribution("2 apples and 3 oranges", 100))
    println(FruitDistribution("100 apples and 1 oranges", 120))
}