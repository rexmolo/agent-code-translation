package main

func IncrList(l []int) []int {
	result := make([]int, len(l))
	for i, e := range l {
		result[i] = e + 1
	}
	return result
}
