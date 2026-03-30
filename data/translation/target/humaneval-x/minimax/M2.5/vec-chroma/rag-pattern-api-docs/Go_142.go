package main

func SumSquares(lst []int) int {
	sum := 0
	for i := 0; i < len(lst); i++ {
		if i%3 == 0 {
			sum += lst[i] * lst[i]
		} else if i%4 == 0 && i%3 != 0 {
			sum += lst[i] * lst[i] * lst[i]
		} else {
			sum += lst[i]
		}
	}
	return sum
}
