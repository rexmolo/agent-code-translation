package main

import "fmt"

func RollingMax(numbers []int) []int {
	result := []int{}
	var runningMax int

	for _, n := range numbers {
		if len(result) == 0 {
			runningMax = n
		} else if n > runningMax {
			runningMax = n
		}
		result = append(result, runningMax)
	}

	return result
}

func main() {
	// Example usage
	numbers := []int{1, 2, 3, 2, 3, 4, 2}
	result := RollingMax(numbers)
	fmt.Println(result)
}