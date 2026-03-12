package main

import "fmt"

func ReverseDelete(s, c string) [2]interface{} {
    // Create a set of characters to delete for O(1) lookup
    deleteSet := make(map[rune]bool)
    for _, ch := range c {
        deleteSet[ch] = true
    }

    // Filter out characters in c
    var result []rune
    for _, ch := range s {
        if !deleteSet[ch] {
            result = append(result, ch)
        }
    }

    resultStr := string(result)

    // Check if palindrome by comparing with reverse
    isPalindrome := true
    runes := []rune(resultStr)
    for i := 0; i < len(runes)/2; i++ {
        if runes[i] != runes[len(runes)-1-i] {
            isPalindrome = false
            break
        }
    }

    return [2]interface{}{resultStr, isPalindrome}
}

func main() {
    fmt.Println(ReverseDelete("abcde", "ae"))
    fmt.Println(ReverseDelete("abcdef", "b"))
    fmt.Println(ReverseDelete("abcdedcba", "ab"))
}
