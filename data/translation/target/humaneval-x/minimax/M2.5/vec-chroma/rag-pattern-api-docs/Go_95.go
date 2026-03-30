package main

import (
	"fmt"
	"unicode"
)

// CheckDictCase returns true if all keys in the dictionary are strings
// in lower case or all keys are strings in upper case, otherwise false.
// Returns false for empty dictionaries.
func CheckDictCase(dict map[interface{}]interface{}) bool {
	// Check if map is empty
	if len(dict) == 0 {
		return false
	}

	state := "start"

	for key := range dict {
		// Check if key is a string
		keyStr, ok := key.(string)
		if !ok {
			return false
		}

		if state == "start" {
			if isAllUpper(keyStr) {
				state = "upper"
			} else if isAllLower(keyStr) {
				state = "lower"
			} else {
				return false
			}
		} else if (state == "upper" && !isAllUpper(keyStr)) || (state == "lower" && !isAllLower(keyStr)) {
			return false
		}
	}

	return state == "upper" || state == "lower"
}

// isAllUpper returns true if all characters in the string are uppercase letters
func isAllUpper(s string) bool {
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

// isAllLower returns true if all characters in the string are lowercase letters
func isAllLower(s string) bool {
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
	testCases := []struct {
		name string
		dict map[interface{}]interface{}
		want bool
	}{
		{"all lowercase", map[interface{}]interface{}{"a": "apple", "b": "banana"}, true},
		{"mixed case keys", map[interface{}]interface{}{"a": "apple", "A": "banana", "B": "banana"}, false},
		{"non-string key", map[interface{}]interface{}{"a": "apple", 8: "banana", "a": "apple"}, false},
		{"mixed case string key", map[interface{}]interface{}{"Name": "John", "Age": "36", "City": "Houston"}, false},
		{"all uppercase", map[interface{}]interface{}{"STATE": "NC", "ZIP": "12345"}, true},
		{"empty dict", map[interface{}]interface{}{}, false},
	}

	for _, tc := range testCases {
		result := CheckDictCase(tc.dict)
		fmt.Printf("%s: got %v, want %v\n", tc.name, result, tc.want)
	}
}
