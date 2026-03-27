package main

import "fmt"

func IsHappy(s string) bool {
    if len(s) < 3 {
        return false
    }

    for i := 0; i <= len(s)-3; i++ {
        if s[i] == s[i+1] || s[i+1] == s[i+2] || s[i] == s[i+2] {
            return false
        }
    }
    return true
}

func main() {
    // Test cases
    fmt.Println(IsHappy("a"))    // false
    fmt.Println(IsHappy("aa"))   // false
    fmt.Println(IsHappy("abcd")) // true
    fmt.Println(IsHappy("aabb")) // false
    fmt.Println(IsHappy("adb"))  // true
    fmt.Println(IsHappy("xyy"))  // false
}