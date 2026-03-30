package main

func GetPositive(l []int) []int {
    var result []int
    for _, v := range l {
        if v > 0 {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    // Example usage
    nums := []int{-1, 2, -4, 5, 6}
    positive := GetPositive(nums)
    _ = positive // Use the result
}
