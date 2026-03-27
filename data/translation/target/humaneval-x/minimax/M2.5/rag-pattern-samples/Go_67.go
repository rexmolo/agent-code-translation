package main

import (
    "strconv"
    "strings"
)

func FruitDistribution(s string, n int) int {
    sum := 0
    tokens := strings.Split(s, " ")
    for _, token := range tokens {
        if num, err := strconv.Atoi(token); err == nil {
            sum += num
        }
    }
    return n - sum
}

func main() {
    // Test examples
    result1 := FruitDistribution("5 apples and 6 oranges", 19)
    println(result1) // Output: 8
    
    result2 := FruitDistribution("0 apples and 1 oranges", 3)
    println(result2) // Output: 2
    
    result3 := FruitDistribution("2 apples and 3 oranges", 100)
    println(result3) // Output: 95
    
    result4 := FruitDistribution("100 apples and 1 oranges", 120)
    println(result4) // Output: 19
}
