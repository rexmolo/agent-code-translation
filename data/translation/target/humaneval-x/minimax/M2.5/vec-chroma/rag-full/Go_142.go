package main

func SumSquares(lst []int) int {
	result := 0
	for i := 0; i < len(lst); i++ {
		if i%3 == 0 {
			// Index is a multiple of 3: square the element
			result += lst[i] * lst[i]
		} else if i%4 == 0 && i%3 != 0 {
			// Index is a multiple of 4 but not 3: cube the element
			result += lst[i] * lst[i] * lst[i]
		} else {
			// Keep the original element
			result += lst[i]
		}
	}
	return result
}