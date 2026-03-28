package main

import "fmt"

func RollingMax(numbers []int) []int {
    if len(numbers) == 0 {
        return []int{}
    }

    runningMax := numbers[0]
    result := make([]int, 0, len(numbers))

    for _, n := range numbers {
        if n > runningMax {
            runningMax = n
        }
        result = append(result, runningMax)
    }

    return result
}

func main() {
    // Test the function
    result := RollingMax([]int{1, 2, 3, 2, 3, 4, 2})
    fmt.Println(result)
}