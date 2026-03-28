package main

import "fmt"

func RollingMax(numbers []int) []int {
	var runningMax int
	hasValue := false
	result := make([]int, 0, len(numbers))

	for _, n := range numbers {
		if !hasValue {
			runningMax = n
			hasValue = true
		} else if n > runningMax {
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