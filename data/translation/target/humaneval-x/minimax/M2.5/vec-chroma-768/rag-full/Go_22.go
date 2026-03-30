package main

import "fmt"

// FilterIntegers filters a slice of interface{} values and returns only those that are integers.
// Example: FilterIntegers([]interface{}{'a', 3.14, 5}) -> []int{5}
func FilterIntegers(values []interface{}) []int {
    result := make([]int, 0, len(values))
    for _, v := range values {
        if i, ok := v.(int); ok {
            result = append(result, i)
        }
    }
    return result
}

func main() {
    // Test the function
    result1 := FilterIntegers([]interface{}{'a', 3.14, 5})
    fmt.Println(result1)

    result2 := FilterIntegers([]interface{}{1, 2, 3, "abc", struct{}{}, []int{}})
    fmt.Println(result2)
}
