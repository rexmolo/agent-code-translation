package main

import (
	"strings"
)

func CheckDictCase(dict map[interface{}]interface{}) bool {
	if len(dict) == 0 {
		return false
	}

	// state: "start", "upper", "lower", "mixed"
	state := "start"

	for key := range dict {
		// Check if key is a string
		strKey, isString := key.(string)
		if !isString {
			return false
		}

		if state == "start" {
			if strKey == strings.ToUpper(strKey) {
				state = "upper"
			} else if strKey == strings.ToLower(strKey) {
				state = "lower"
			} else {
				// Key has mixed case (e.g., "John", "aBc")
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
