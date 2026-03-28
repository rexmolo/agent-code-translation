func RollingMax(numbers []int) []int {
	result := make([]int, 0, len(numbers))

	for _, n := range numbers {
		if len(result) == 0 || n > result[len(result)-1] {
			result = append(result, n)
		} else {
			result = append(result, result[len(result)-1])
		}
	}

	return result
}