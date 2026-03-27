package main

import "strings"

func FilterBySubstring(stringList []string, substring string) []string {
    var result []string
    for _, s := range stringList {
        if strings.Contains(s, substring) {
            result = append(result, s)
        }
    }
    return result
}

func main() {
    // Test cases
    test1 := FilterBySubstring([]string{}, "a")
    println("Test 1:", test1)
    
    test2 := FilterBySubstring([]string{"abc", "bacd", "cde", "array"}, "a")
    println("Test 2:", test2)
}