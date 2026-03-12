func GetPositive(l []int) []int {
    result := make([]int, 0)
    for _, e := range l {
        if e > 0 {
            result = append(result, e)
        }
    }
    return result
}