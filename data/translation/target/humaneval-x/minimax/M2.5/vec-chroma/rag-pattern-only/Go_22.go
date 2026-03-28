package main

func FilterIntegers(values []interface{}) []int {
	result := make([]int, 0, len(values))
	for _, v := range values {
		if intVal, ok := v.(int); ok {
			result = append(result, intVal)
		}
	}
	return result
}
