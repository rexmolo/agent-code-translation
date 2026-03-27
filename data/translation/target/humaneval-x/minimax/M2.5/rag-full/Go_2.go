package main

import (
    "fmt"
    "math"
)

func TruncateNumber(number float64) float64 {
    // Return the decimal (fractional) part of the number
    // In Go, the equivalent of Python's % for floats is math.Mod
    return math.Mod(number, 1)
}

func main() {
    // Example usage
    fmt.Println(TruncateNumber(3.5))
}