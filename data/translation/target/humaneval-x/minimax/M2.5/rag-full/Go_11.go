package main

import "strings"

func StringXor(a string, b string) string {
	result := make([]byte, 0, len(a))

	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			result = append(result, '0')
		} else {
			result = append(result, '1')
		}
	}

	return string(result)
}

func main() {
	// Example usage
	// fmt.Println(StringXor("010", "110")) // Output: "100"
}
