package main

import "fmt"

func IsEqualToSumEven(n int) bool {
    return n%2 == 0 && n >= 8
}

func main() {
    // Test examples from the docstring
    fmt.Println(IsEqualToSumEven(4)) // false
    fmt.Println(IsEqualToSumEven(6)) // false
    fmt.Println(IsEqualToSumEven(8)) // true
}
