package main

func CheckDictCase(dict map[interface{}]interface{}) bool {
    if len(dict) == 0 {
        return false
    }

    state := "start"
    for key := range dict {
        if _, ok := key.(string); !ok {
            return false
        }

        s := key.(string)
        if state == "start" {
            if isUpper(s) {
                state = "upper"
            } else if isLower(s) {
                state = "lower"
            } else {
                break
            }
        } else if (state == "upper" && !isUpper(s)) || (state == "lower" && !isLower(s)) {
            state = "mixed"
            break
        } else {
            break
        }
    }

    return state == "upper" || state == "lower"
}

func isUpper(s string) bool {
    if len(s) == 0 {
        return false
    }
    for _, c := range s {
        if c < 'A' || c > 'Z' {
            return false
        }
    }
    return true
}

func isLower(s string) bool {
    if len(s) == 0 {
        return false
    }
    for _, c := range s {
        if c < 'a' || c > 'z' {
            return false
        }
    }
    return true
}
