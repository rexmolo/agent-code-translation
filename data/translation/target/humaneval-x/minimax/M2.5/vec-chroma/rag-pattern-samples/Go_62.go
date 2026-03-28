package main

func Derivative(xs []int) []int {
	var result []int
	for i, x := range xs {
		result = append(result, i*x)
	}
	return result[1:]
}