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
				return false
			}
		} else if (state == "upper" && strKey != strings.ToUpper(strKey)) || (state == "lower" && strKey != strings.ToLower(strKey)) {
			return false
		} else {
			break
		}
	}

	return state == "upper" || state == "lower"
}

func main() {
	// Test cases
	testCases := []map[interface{}]interface{}{
		{"a": "apple", "b": "banana"},                           // true
		{"a": "apple", "A": "banana", "B": "banana"},        // false
		{"a": "apple", 8: "banana", "a": "apple"},            // false
		{"Name": "John", "Age": "36", "City": "Houston"},    // false
		{"STATE": "NC", "ZIP": "12345"},                       // true
	}

	for i, tc := range testCases {
		result := CheckDictCase(tc)
		println("Test", i+1, ":", result)
	}
}
