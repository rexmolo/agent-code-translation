package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func ParseNestedParens(parenString string) []int {
    groups := strings.Split(parenString, " ")
    
    var result []int
    
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
        
        result = append(result, maxDepth)
    }
    
    return result
}

func main() {
    // Read from stdin and write to stdout
    scanner := bufio.NewScanner(os.Stdin)
    
    var inputs []string
    for scanner.Scan() {
        inputs = append(inputs, scanner.Text())
    }
    
    for _, input := range inputs {
        result := ParseNestedParens(input)
        fmt.Println(result)
    }
}