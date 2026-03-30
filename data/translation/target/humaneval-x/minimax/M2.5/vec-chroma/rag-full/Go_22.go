package main

// FilterIntegers filters the given slice of interface{} values and returns only integers.
func FilterIntegers(values []interface{}) []int {
	result := make([]int, 0, len(values))
	for _, v := range values {
		if intVal, ok := v.(int); ok {
			result = append(result, intVal)
		}
	}
	return result
}

func main() {
	// Example usage
	values := []interface{}{"a", 3.14, 5}
	filtered := FilterIntegers(values)
	_ = filtered // Use the filtered result
}
