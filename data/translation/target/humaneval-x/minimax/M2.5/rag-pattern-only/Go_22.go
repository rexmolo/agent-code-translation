func FilterIntegers(values []interface{}) []int {
    var result []int
    for _, v := range values {
        if i, ok := v.(int); ok {
            result = append(result, i)
        }
    }
    return result
}