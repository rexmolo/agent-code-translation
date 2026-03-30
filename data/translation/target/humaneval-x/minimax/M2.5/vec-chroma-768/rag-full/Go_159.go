package main

func Eat(number, need, remaining int) []int {
	if need <= remaining {
		return []int{number + need, remaining - need}
	}
	return []int{number + remaining, 0}
}

func main() {
	// Test cases for verification
	testCases := []struct {
		number, need, remaining int
	}{
		{5, 6, 10},
		{4, 8, 9},
		{1, 10, 10},
		{2, 11, 5},
	}
	for _, tc := range testCases {
		result := Eat(tc.number, tc.need, tc.remaining)
		_ = result // Use result to avoid unused variable error
	}
}
