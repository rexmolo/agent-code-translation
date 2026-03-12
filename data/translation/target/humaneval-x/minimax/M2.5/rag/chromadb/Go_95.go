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
		str, ok := key.(string)
		if !ok {
			return false
		}

		if state == "start" {
			if str == strings.ToUpper(str) {
				state = "upper"
			} else if str == strings.ToLower(str) {
				state = "lower"
			} else {
				// key is neither all upper nor all lower (e.g., "John")
				return false
			}
		} else if state == "upper" {
			if str != strings.ToUpper(str) {
				return false
			}
		} else if state == "lower" {
			if str != strings.ToLower(str) {
				return false
			}
		}
	}

	return state == "upper" || state == "lower"
}