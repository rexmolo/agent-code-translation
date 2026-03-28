package main

func IncrList(l []int) []int {
    result := make([]int, 0, len(l))
    for _, v := range l {
        result = append(result, v+1)
    }
    return result
}