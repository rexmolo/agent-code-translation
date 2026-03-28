package main

import "fmt"

func GreatestCommonDivisor(a int, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func main() {
    fmt.Println(GreatestCommonDivisor(3, 5))  // Output: 1
    fmt.Println(GreatestCommonDivisor(25, 15)) // Output: 5
}