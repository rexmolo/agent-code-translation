package main

import "fmt"

// Add adds two numbers x and y.
func Add(x int, y int) int {
    return x + y
}

func main() {
    // Example usage from docstring
    fmt.Println(Add(2, 3)) // 5
    fmt.Println(Add(5, 7)) // 12
}