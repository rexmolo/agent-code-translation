package main

import "fmt"

func FilterIntegers(values []interface{}) []int {
    result := []int{}
    for _, v := range values {
        if intVal, ok := v.(int); ok {
            result = append(result, intVal)
        }
    }
    return result
}

func main() {
    // Example usage
    fmt.Println(FilterIntegers([]interface{}{"a", 3.14, 5}))
    fmt.Println(FilterIntegers([]interface{}{1, 2, 3, "abc", struct{}{}, []interface{}{}}))
}