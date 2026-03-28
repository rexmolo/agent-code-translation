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
	for key, _ := range dict {
		str, ok := key.(string)
		if !ok {
			state = "mixed"
			break
		}

		if state == "start" {
			if strings.ToUpper(str) == str {
				state = "upper"
			} else if strings.ToLower(str) == str {
				state = "lower"
			} else {
				state = "mixed"
				break
			}
		} else if (state == "upper" && strings.ToUpper(str) != str) || (state == "lower" && strings.ToLower(str) != str) {
			state = "mixed"
			break
		} else {
			state = "mixed"
			break
		}
	}
	return state == "upper" || state == "lower"
}

func main() {
	// Test cases
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"a": "apple", "b": "banana"}))           // true
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"a": "apple", "A": "banana", "B": "banana"})) // false
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"a": "apple", 8: "banana", "a": "apple"}))    // false
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"Name": "John", "Age": "36", "City": "Houston"})) // false
	fmt.Println(CheckDictCase(map[interface{}]interface{}{"STATE": "NC", "ZIP": "12345"}))            // true
	fmt.Println(CheckDictCase(map[interface{}]interface{}{})) // false
}
