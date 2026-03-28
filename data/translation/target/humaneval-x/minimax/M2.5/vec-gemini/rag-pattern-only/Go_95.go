package main

import (
	"strings"
)

func CheckDictCase(dict map[interface{}]interface{}) bool {
	if len(dict) == 0 {
		return false
	}

	state := "start"

	for key := range dict {
		// Check if key is a string
		strKey, ok := key.(string)
		if !ok {
			return false
		}

		if state == "start" {
			if strKey == strings.ToUpper(strKey) {
				state = "upper"
			} else if strKey == strings.ToLower(strKey) {
				state = "lower"
			} else {
				// Key is neither all upper nor all lower (e.g., "Ab")
				return false
			}
		} else if state == "upper" {
			if strKey != strings.ToUpper(strKey) {
				return false
			}
		} else if state == "lower" {
			if strKey != strings.ToLower(strKey) {
				return false
			}
		}
	}

	return state == "upper" || state == "lower"
}

func main() {
	// Test cases
	testCases := []map[interface{}]interface{}{
		{"a": "apple", "b": "banana"},
		{"a": "apple", "A": "banana", "B": "banana"},
		{"a": "apple", 8: "banana", "a": "apple"},
		{"Name": "John", "Age": "36", "City": "Houston"},
		{"STATE": "NC", "ZIP": "12345"},
	}

	for i, tc := range testCases {
		result := CheckDictCase(tc)
		.Printf("Test %d: %v\n", i+1, result)
	}
}
