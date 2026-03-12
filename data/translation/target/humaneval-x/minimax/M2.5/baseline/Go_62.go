package main

func Derivative(xs []int) []int {
	result := make([]int, len(xs)-1)
	for i, x := range xs {
		if i > 0 {
			result[i-1] = i * x
		}
	}
	return result
}

func main() {
	// Example usage
	println(Derivative([]int{3, 1, 2, 4, 5})) // Would print slice address, use fmt for human readable
}