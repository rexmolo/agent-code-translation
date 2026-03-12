package main

import (
	"fmt"
	"unicode"
)

// isUpper checks if a string contains at least one cased character
// and all cased characters are uppercase. This mimics Python's str.isupper().
func isUpper(s string) bool {
	hasCased := false
	for _, r := range s {
		if unicode.IsUpper(r) {
			hasCased = true
		} else if unicode.IsLower(r) {
			return false // Contains a lowercase letter
		}
	}
	return hasCased
}

// isLower checks if a string contains at least one cased character
// and all cased characters are lowercase. This mimics Python's str.islower().
func isLower(s string) bool {
	hasCased := false
	for _, r := range s {
		if unicode.IsLower(r) {
			hasCased = true
		} else if unicode.IsUpper(r) {
			return false // Contains an uppercase letter
		}
	}
	return hasCased
}

// CheckDictCase translates the Python function check_dict_case.
// Given a map, it returns true if all keys are strings in lower
// case or all keys are strings in upper case, otherwise it returns false.
// It returns false for an empty map.
func CheckDictCase(dict map[interface{}]interface{}) bool {
	if len(dict) == 0 {
		return false
	}

	state := "start"

	for key := range dict {
		keyStr, ok := key.(string)
		if !ok {
			state = "mixed"
			break
		}

		if state == "start" {
			if isUpper(keyStr) {
				state = "upper"
			} else if isLower(keyStr) {
				state = "lower"
			} else {
				// First key is mixed case or has no cased characters.
				state = "mixed"
				break
			}
		} else if (state == "upper" && !isUpper(keyStr)) || (state == "lower" && !isLower(keyStr)) {
			// Key case does not match the established state
			state = "mixed"
			break
		} else {
			// This branch mimics the 'else: break' in the original Python code.
			// It is reached if a key's case *matches* the established state
			// (on the 2nd iteration onwards), causing the loop to exit prematurely.
			break
		}
	}

	return state == "upper" || state == "lower"
}

// main function to demonstrate CheckDictCase with the provided examples.
func main() {
	fmt.Println("Running examples:")

	// Example 1: {"a":"apple", "b":"banana"} -> True
	dict1 := map[interface{}]interface{}{"a": "apple", "b": "banana"}
	fmt.Printf("check_dict_case({\"a\":\"apple\", \"b\":\"banana\"}) -> %t\n", CheckDictCase(dict1))

	// Example 2: {"a":"apple", "A":"banana", "B":"banana"} -> False
	dict2 := map[interface{}]interface{}{"a": "apple", "A": "banana", "B": "banana"}
	fmt.Printf("check_dict_case({\"a\":\"apple\", \"A\":\"banana\", ...}) -> %t\n", CheckDictCase(dict2))

	// Example 3: {"a":"apple", 8:"banana", ...} -> False
	dict3 := map[interface{}]interface{}{"a": "apple", 8: "banana"}
	fmt.Printf("check_dict_case({\"a\":\"apple\", 8:\"banana\", ...}) -> %t\n", CheckDictCase(dict3))

	// Example 4: {"Name":"John", "Age":"36", ...} -> False
	dict4 := map[interface{}]interface{}{"Name": "John", "Age": "36", "City": "Houston"}
	fmt.Printf("check_dict_case({\"Name\":\"John\", \"Age\":\"36\", ...}) -> %t\n", CheckDictCase(dict4))

	// Example 5: {"STATE":"NC", "ZIP":"12345" } -> True
	dict5 := map[interface{}]interface{}{"STATE": "NC", "ZIP": "12345"}
	fmt.Printf("check_dict_case({\"STATE\":\"NC\", \"ZIP\":\"12345\"}) -> %t\n", CheckDictCase(dict5))

	// Empty dictionary -> False
	dict6 := map[interface{}]interface{}{}
	fmt.Printf("check_dict_case({}) -> %t\n", CheckDictCase(dict6))
}
