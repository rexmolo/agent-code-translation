package main

func SumSquares(lst []int) int {
	result := 0
	for i := 0; i < len(lst); i++ {
		if i%3 == 0 {
			result += lst[i] * lst[i]
		} else if i%4 == 0 && i%3 != 0 {
			result += lst[i] * lst[i] * lst[i]
		} else {
			result += lst[i]
		}
	}
	return result
}