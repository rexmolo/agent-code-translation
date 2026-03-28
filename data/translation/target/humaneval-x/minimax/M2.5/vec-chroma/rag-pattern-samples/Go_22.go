package main

func FilterIntegers(values []interface{}) []int {
    var result []int
    for _, v := range values {
        if _, ok := v.(int); ok {
            result = append(result, v.(int))
        }
    }
    return result
}