package main

func SumSquares(lst []int) int {
	sum := 0
	for i := 0; i < len(lst); i++ {
		val := lst[i]
		if i%3 == 0 {
			// Index is a multiple of 3: square the element
			sum += val * val
		} else if i%4 == 0 {
			// Index is a multiple of 4 but NOT multiple of 3: cube the element
			sum += val * val * val
		} else {
			// Index is neither multiple of 3 nor 4: keep the element as is
			sum += val
		}
	}
	return sum
}