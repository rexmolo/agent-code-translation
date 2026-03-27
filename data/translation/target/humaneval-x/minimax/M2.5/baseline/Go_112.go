package main

import "fmt"

func ReverseDelete(s, c string) [2]interface{} {
    // Build a map of characters to delete for efficient lookup
    deleteChars := make(map[rune]bool)
    for _, char := range c {
        deleteChars[char] = true
    }
    
    // Filter out characters that are in c
    var result string
    for _, char := range s {
        if !deleteChars[char] {
            result += string(char)
        }
    }
    
    // Check if palindrome
    isPalindrome := true
    n := len(result)
    for i := 0; i < n/2; i++ {
        if result[i] != result[n-1-i] {
            isPalindrome = false
            break
        }
    }
    
    return [2]interface{}{result, isPalindrome}
}

func main() {
    // Test cases
    fmt.Println(ReverseDelete("abcde", "ae"))     // [bcd false]
    fmt.Println(ReverseDelete("abcdef", "b"))     // [acdef false]
    fmt.Println(ReverseDelete("abcdedcba", "ab")) // [cdedc true]
}