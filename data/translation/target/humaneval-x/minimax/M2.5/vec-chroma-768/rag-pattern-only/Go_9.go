package main

import "fmt"

func RollingMax(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	result := make([]int, 0, len(numbers))
	runningMax := numbers[0]

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
	numbers := []int{1, 2, 3, 2, 3, 4, 2}
	result := RollingMax(numbers)
	fmt.Println(result)
}
