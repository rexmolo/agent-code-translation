package main

func Derivative(xs []int) []int {
	result := make([]int, 0, len(xs)-1)
	for i, x := range xs {
		if i > 0 {
			result = append(result, i*x)
		}
	}
	return result
}