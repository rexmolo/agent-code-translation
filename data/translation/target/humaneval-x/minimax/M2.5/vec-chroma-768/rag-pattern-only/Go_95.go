package main

import (
	"fmt"
	"unicode"
)

func CheckDictCase(dict map[interface{}]interface{}) bool {
	// Return False if the dictionary is empty
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
			if isUpperCase(strKey) {
				state = "upper"
			} else if isLowerCase(strKey) {
				state = "lower"
			} else {
				// Key is neither all uppercase nor all lowercase
				return false
			}
		} else if state == "upper" && !isUpperCase(strKey) {
			return false
		} else if state == "lower" && !isLowerCase(strKey) {
			return false
		}
	}

	return state == "upper" || state == "lower"
}

// isUpperCase checks if all characters in the string are uppercase letters
func isUpperCase(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// isLowerCase checks if all characters in the string are lowercase letters
func isLowerCase(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	testCases := []map[interface{}]interface{}{
		{"a": "apple", "b": "banana"},                          // true
		{"a": "apple", "A": "banana", "B": "banana"},        // false
		{"a": "apple", 8: "banana", "a": "apple"},            // false
		{"Name": "John", "Age": "36", "City": "Houston"},   // false
		{"STATE": "NC", "ZIP": "12345"},                       // true
		{},                                                         // false (empty)
	}

	expected := []bool{true, false, false, false, true, false}

	for i, tc := range testCases {
		result := CheckDictCase(tc)
		fmt.Printf("Test %d: %v (expected: %v)\n", i+1, result, expected[i])
	}
}
