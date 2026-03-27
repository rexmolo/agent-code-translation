package main

func GetPositive(l []int) []int {
    result := make([]int, 0, len(l))
    for _, v := range l {
        if v > 0 {
            result = append(result, v)
        }
    }
    return result
}