package main

func RollingMax(numbers []int) []int {
    if len(numbers) == 0 {
        return []int{}
    }

    runningMax := numbers[0]
    result := make([]int, 0, len(numbers))
    result = append(result, runningMax)

    for i := 1; i < len(numbers); i++ {
        if numbers[i] > runningMax {
            runningMax = numbers[i]
        }
        result = append(result, runningMax)
    }

    return result
}
