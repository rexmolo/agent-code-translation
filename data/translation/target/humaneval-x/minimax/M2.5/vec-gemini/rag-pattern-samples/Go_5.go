package main

import "fmt"

func Intersperse(numbers []int, delimeter int) []int {
    if len(numbers) == 0 {
        return []int{}
    }

    // Pre-allocate result with capacity to avoid reallocations
    // For n elements, we need 2*n - 1 positions (n elements + n-1 delimiters)
    result := make([]int, 0, 2*len(numbers)-1)

    // Append all elements except the last, with delimiter after each
    for i := 0; i < len(numbers)-1; i++ {
        result = append(result, numbers[i])
        result = append(result, delimeter)
    }

    // Append the last element without a trailing delimiter
    result = append(result, numbers[len(numbers)-1])

    return result
}

func main() {
    // Test cases from docstring
    fmt.Println(Intersperse([]int{}, 4))
    fmt.Println(Intersperse([]int{1, 2, 3}, 4))
}
