package main

import (
    "fmt"
    "strings"
)

func ReverseDelete(s, c string) [2]interface{} {
    // Filter out characters that are in string c
    var result strings.Builder
    for _, char := range s {
        if !strings.Contains(c, string(char)) {
            result.WriteRune(char)
        }
    }
    
    resultStr := result.String()
    
    // Check if palindrome by comparing with reversed string
    isPalindrome := resultStr == reverseString(resultStr)
    
    return [2]interface{}{resultStr, isPalindrome}
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    // Test cases
    fmt.Println(ReverseDelete("abcde", "ae"))      // [bcd false]
    fmt.Println(ReverseDelete("abcdef", "b"))     // [acdef false]
    fmt.Println(ReverseDelete("abcdedcba", "ab")) // [cdedc true]
}
