package main

import "fmt"

func isPalindrome(str string) bool {
    // Compare string with its reverse
    return str == reverseString(str)
}

func reverseString(str string) string {
    runes := []rune(str)
    for i, j := 0; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func MakePalindrome(str string) string {
    if str == "" {
        return ""
    }

    beginningOfSuffix := 0

    for !isPalindrome(str[beginningOfSuffix:]) {
        beginningOfSuffix++
    }

    prefix := str[:beginningOfSuffix]
    return str + reverseString(prefix)
}

func main() {
    // Test cases from docstring
    fmt.Println(MakePalindrome(""))    // Expected: ""
    fmt.Println(MakePalindrome("cat")) // Expected: "catac"
    fmt.Println(MakePalindrome("cata")) // Expected: "catac"
}
