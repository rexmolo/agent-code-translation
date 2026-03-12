package main

import "fmt"

func RollingMax(numbers []int) []int {
	runningMax := 0
	found := false
	result := make([]int, 0, len(numbers))

	for _, n := range numbers {
		if !found {
			runningMax = n
			found = true
		} else if n > runningMax {
			runningMax = n
		}

		result = append(result, runningMax)
	}

	return result
}

func main() {
	fmt.Println(RollingMax([]int{1, 2, 3, 2, 3, 4, 2}))
}
