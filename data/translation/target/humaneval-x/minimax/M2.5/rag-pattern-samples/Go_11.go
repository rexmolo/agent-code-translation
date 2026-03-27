package main

import "fmt"

func StringXor(a string, b string) string {
	xor := func(i, j byte) byte {
		if i == j {
			return '0'
		}
		return '1'
	}

	var result []byte
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		result = append(result, xor(a[i], b[i]))
	}

	return string(result)
}

func main() {
	fmt.Println(StringXor("010", "110"))
}