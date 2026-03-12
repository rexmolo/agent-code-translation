func GetPositive(l []int) []int {
    result := make([]int, 0)
    for _, v := range l {
        if v > 0 {
            result = append(result, v)
        }
    }
    return result
}