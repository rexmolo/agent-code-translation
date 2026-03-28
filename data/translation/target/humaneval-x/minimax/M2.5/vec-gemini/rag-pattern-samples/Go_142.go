func SumSquares(lst []int) int {
	result := 0
	for i, v := range lst {
		if i%3 == 0 {
			result += v * v
		} else if i%4 == 0 && i%3 != 0 {
			result += v * v * v
		} else {
			result += v
		}
	}
	return result
}