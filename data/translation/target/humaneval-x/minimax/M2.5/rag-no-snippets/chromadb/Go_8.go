package main

func SumProduct(numbers []int) [2]int {
	sumValue := 0
	prodValue := 1

	for _, n := range numbers {
		sumValue += n
		prodValue *= n
	}
	return [2]int{sumValue, prodValue}
}