package main

import "fmt"

func Intersperse(numbers []int, delimeter int) []int {
    if len(numbers) == 0 {
        return []int{}
    }

    result := make([]int, 0, len(numbers)*2-1)

    for i := 0; i < len(numbers)-1; i++ {
        result = append(result, numbers[i])
        result = append(result, delimeter)
    }

    result = append(result, numbers[len(numbers)-1])

    return result
}

func main() {
    // Test cases
    fmt.Println(Intersperse([]int{}, 4))      // []
    fmt.Println(Intersperse([]int{1, 2, 3}, 4)) // [1 4 2 4 3]
}