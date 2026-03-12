package main

import (
	"fmt"
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
			if strings.ToUpper(strKey) == strKey {
				state = "upper"
			} else if strings.ToLower(strKey) == strKey {
				state = "lower"
			} else {
				// key is neither all uppercase nor all lowercase
				return false
			}
		} else if state == "upper" {
			if strings.ToUpper(strKey) != strKey {
				return false
			}
		} else if state == "lower" {
			if strings.ToLower(strKey) != strKey {
				return false
			}
		}
	}

	return state == "upper" || state == "lower"
}

func main() {
	// Test cases
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"a": "apple", "b": "banana"}))                              // true
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"a": "apple", "A": "banana", "B": "banana"}))              // false
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"a": "apple", 8: "banana", "a": "apple"}))                 // false
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"Name": "John", "Age": "36", "City": "Houston"}))          // false
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"STATE": "NC", "ZIP": "12345"}))                             // true
	fmt.Println(CheckDictCase(map[interface{}]interface{}{}))                                                             // false
}
