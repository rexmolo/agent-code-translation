package main

// RollingMax from a given list of integers, generates a list of rolling maximum element found until given moment
// in the sequence.
func RollingMax(numbers []int) []int {
	result := make([]int, 0, len(numbers))
	if len(numbers) == 0 {
		return result
	}

	var runningMax int
	for i, n := range numbers {
		if i == 0 {
			runningMax = n
		} else if n > runningMax {
			runningMax = n
		}
		result = append(result, runningMax)
	}

	return result
}
