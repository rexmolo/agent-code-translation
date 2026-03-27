package main

func FilterIntegers(values []interface{}) []int {
    result := make([]int, 0)
    for _, v := range values {
        if n, ok := v.(int); ok {
            result = append(result, n)
        }
    }
    return result
}