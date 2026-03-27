package main

func SumSquares(lst []int) int {
	result := make([]int, 0, len(lst))
	for i := 0; i < len(lst); i++ {
		if i%3 == 0 {
			// Index is a multiple of 3: square the element
			result = append(result, lst[i]*lst[i])
		} else if i%4 == 0 && i%3 != 0 {
			// Index is a multiple of 4 but not 3: cube the element
			result = append(result, lst[i]*lst[i]*lst[i])
		} else {
			// Index is not a multiple of 3 or 4: keep original
			result = append(result, lst[i])
		}
	}

	// Sum all elements in result
	sum := 0
	for _, v := range result {
		sum += v
	}
	return sum
}
