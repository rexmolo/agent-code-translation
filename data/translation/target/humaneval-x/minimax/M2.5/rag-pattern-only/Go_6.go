package main

import (
    "fmt"
    "strings"
)

func ParseNestedParens(parenString string) []int {
    groups := strings.Split(parenString, " ")

    var results []int

    for _, group := range groups {
        if group == "" {
            continue
        }

        depth := 0
        maxDepth := 0
        for _, c := range group {
            if c == '(' {
                depth++
                if depth > maxDepth {
                    maxDepth = depth
                }
            } else {
                depth--
            }
        }

        results = append(results, maxDepth)
    }

    return results
}

func main() {
    // Test the function
    input := "(()()) ((())) () ((())()')"
    result := ParseNestedParens(input)
    fmt.Println(result)
}
