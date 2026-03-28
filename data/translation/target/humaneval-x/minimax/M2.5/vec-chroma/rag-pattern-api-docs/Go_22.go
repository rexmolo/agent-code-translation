package main

import "fmt"

func FilterIntegers(values []interface{}) []int {
    result := make([]int, 0, len(values))
    for _, v := range values {
        if n, ok := v.(int); ok {
            result = append(result, n)
        }
    }
    return result
}

func main() {
    // Test the function with the examples from the docstring
    fmt.Println(FilterIntegers([]interface{}{"a", 3.14, 5}))
    fmt.Println(FilterIntegers([]interface{}{1, 2, 3, "abc", map[string]interface{}{}, []interface{}{}}))
}
