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
    // Test cases from docstring
    result1 := FilterBySubstring([]string{}, "a")
    println(result1) // []
    
    result2 := FilterBySubstring([]string{"abc", "bacd", "cde", "array"}, "a")
    println(result2) // [abc bacd array]
}