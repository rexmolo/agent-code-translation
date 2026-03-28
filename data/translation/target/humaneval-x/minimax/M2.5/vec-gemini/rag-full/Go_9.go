package main

func RollingMax(numbers []int) []int {
	if len(numbers) == 0 {
		return nil
	}

	runningMax := numbers[0]
	result := make([]int, 0, len(numbers))
	result = append(result, runningMax)

	for _, n := range numbers[1:] {
		if n > runningMax {
			runningMax = n
		}
		result = append(result, runningMax)
	}

	return result
}
