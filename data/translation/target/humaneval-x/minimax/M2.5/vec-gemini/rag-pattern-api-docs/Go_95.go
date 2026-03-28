package main

import "strings"

func CheckDictCase(dict map[interface{}]interface{}) bool {
    // Return False if dictionary is empty
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
                // Key is neither all uppercase nor all lowercase (e.g., "Name")
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

// Optional: Add main function for testing
func main() {
    // Test cases
    test1 := map[interface{}]interface{}{"a": "apple", "b": "banana"}
    test2 := map[interface{}]interface{}{"a": "apple", "A": "banana", "B": "banana"}
    test3 := map[interface{}]interface{}{"a": "apple", 8: "banana", "a": "apple"}
    test4 := map[interface{}]interface{}{"Name": "John", "Age": "36", "City": "Houston"}
    test5 := map[interface{}]interface{}{"STATE": "NC", "ZIP": "12345"}

    println(CheckDictCase(test1)) // true
    println(CheckDictCase(test2)) // false
    println(CheckDictCase(test3)) // false
    println(CheckDictCase(test4)) // false
    println(CheckDictCase(test5)) // true
}
