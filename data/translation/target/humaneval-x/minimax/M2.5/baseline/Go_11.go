package main

func StringXor(a string, b string) string {
	// Create a byte slice to store the result
	result := make([]byte, len(a))

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	return string(result)
}