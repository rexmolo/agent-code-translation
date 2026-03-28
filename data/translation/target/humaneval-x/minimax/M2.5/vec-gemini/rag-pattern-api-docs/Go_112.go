package main

import (
    "strings"
)

func ReverseDelete(s, c string) [2]interface{} {
    // Create a map of characters in string c for O(1) lookup
    charMap := make(map[rune]bool)
    for _, char := range c {
        charMap[char] = true
    }
    
    // Filter characters: keep only those not in c
    var result []rune
    for _, char := range s {
        if !charMap[char] {
            result = append(result, char)
        }
    }
    
    resultStr := string(result)
    
    // Check if the result is a palindrome
    isPalindrome := resultStr == reverseString(resultStr)
    
    return [2]interface{}{resultStr, isPalindrome}
}

// Helper function to reverse a string
func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    // Test the function
    result1 := ReverseDelete("abcde", "ae")
    println(result1[0], result1[1]) // Expected: bcd false
    
    result2 := ReverseDelete("abcdef", "b")
    println(result2[0], result2[1]) // Expected: acdef false
    
    result3 := ReverseDelete("abcdedcba", "ab")
    println(result3[0], result3[1]) // Expected: cdedc true
}
