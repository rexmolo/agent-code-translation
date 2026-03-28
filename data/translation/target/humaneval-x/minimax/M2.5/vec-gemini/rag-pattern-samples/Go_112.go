package main

import "fmt"

func ReverseDelete(s, c string) [2]interface{} {
    // Create a map for quick lookup of characters to delete
    charMap := make(map[rune]bool)
    for _, char := range c {
        charMap[char] = true
    }

    // Filter out characters that are in c (equivalent to Python list comprehension)
    var result string
    for _, char := range s {
        if !charMap[char] {
            result += string(char)
        }
    }

    // Check if palindrome by comparing with reverse
    isPalindrome := true
    runes := []rune(result)
    for i := 0; i < len(runes)/2; i++ {
        if runes[i] != runes[len(runes)-1-i] {
            isPalindrome = false
            break
        }
    }

    return [2]interface{}{result, isPalindrome}
}

func main() {
    // Test cases
    result1 := ReverseDelete("abcde", "ae")
    fmt.Printf("%v\n", result1) // [bcd false]

    result2 := ReverseDelete("abcdef", "b")
    fmt.Printf("%v\n", result2) // [acdef false]

    result3 := ReverseDelete("abcdedcba", "ab")
    fmt.Printf("%v\n", result3) // [cdedc true]
}