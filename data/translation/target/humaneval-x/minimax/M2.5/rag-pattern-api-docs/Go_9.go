package main

import "fmt"

func RollingMax(numbers []int) []int {
    if len(numbers) == 0 {
        return []int{}
    }

    runningMax := numbers[0]
    result := make([]int, 0, len(numbers))
    result = append(result, runningMax)

    for i := 1; i < len(numbers); i++ {
        n := numbers[i]
        if n > runningMax {
            runningMax = n
        }
        result = append(result, runningMax)
    }

    return result
}

func main() {
    // Test examples
    fmt.Println(RollingMax([]int{1, 2, 3, 2, 3, 4, 2}))
    fmt.Println(RollingMax([]int{}))
    fmt.Println(RollingMax([]int{5}))
}
