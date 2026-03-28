package main

import (
    "fmt"
)

func SeparateParenGroups(parenString string) []string {
    result := []string{}
    currentString := []byte{}
    currentDepth := 0

    for _, c := range parenString {
        if c == '(' {
            currentDepth++
            currentString = append(currentString, c)
        } else if c == ')' {
            currentDepth--
            currentString = append(currentString, c)

            if currentDepth == 0 {
                result = append(result, string(currentString))
                currentString = currentString[:0]
            }
        }
    }

    return result
}

func main() {
    // Test the function
    input := "( ) (( )) (( )( ))"
    result := SeparateParenGroups(input)
    fmt.Println(result)
    // Output: [() (()) (()())]
}
