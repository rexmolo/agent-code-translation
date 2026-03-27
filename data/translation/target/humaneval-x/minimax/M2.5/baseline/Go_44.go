package main

import "fmt"

func ChangeBase(x int, base int) string {
    if x == 0 {
        return ""
    }
    
    result := ""
    for x > 0 {
        digit := x % base
        result = fmt.Sprintf("%d", digit) + result
        x /= base
    }
    return result
}

func main() {
    fmt.Println(ChangeBase(8, 3)) // Output: 22
    fmt.Println(ChangeBase(8, 2)) // Output: 1000
    fmt.Println(ChangeBase(7, 2)) // Output: 111
}