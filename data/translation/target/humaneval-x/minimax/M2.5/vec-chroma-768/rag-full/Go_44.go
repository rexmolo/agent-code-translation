package main

import (
    "strconv"
)

func ChangeBase(x int, base int) string {
    if x == 0 {
        return "0"
    }
    
    result := ""
    for x > 0 {
        result = strconv.Itoa(x%base) + result
        x /= base
    }
    return result
}

func main() {
    // Test cases
    println(ChangeBase(8, 3))  // Expected: 22
    println(ChangeBase(8, 2))  // Expected: 1000
    println(ChangeBase(7, 2))  // Expected: 111
}