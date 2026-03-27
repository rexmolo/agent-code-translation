package main

func StringXor(a string, b string) string {
	aBytes := []byte(a)
	bBytes := []byte(b)

	// Use the length of the shorter string (similar to Python's zip)
	n := len(aBytes)
	if len(bBytes) < n {
		n = len(bBytes)
	}

	result := make([]byte, n)
	for i := 0; i < n; i++ {
		if aBytes[i] == bBytes[i] {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	return string(result)
}

func main() {
	// Test examples
	println(StringXor("010", "110")) // Expected: "100"
}