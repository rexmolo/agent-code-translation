package main

import "fmt"

func RollingMax(numbers []int) []int {
    if len(numbers) == 0 {
        return []int{}
    }

    result := make([]int, len(numbers))
    runningMax := numbers[0]
    result[0] = runningMax

    for i := 1; i < len(numbers); i++ {
        if numbers[i] > runningMax {
            runningMax = numbers[i]
        }
        result[i] = runningMax
    }

    return result
}

func main() {
    // Test the function
    fmt.Println(RollingMax([]int{1, 2, 3, 2, 3, 4, 2}))
}
