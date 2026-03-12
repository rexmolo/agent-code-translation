package main

import (
    "strconv"
    "strings"
)

func CompareOne(a, b interface{}) interface{} {
    // Convert a to float64 for comparison
    floatA, errA := toFloat64(a)
    if errA != nil {
        return nil
    }
    
    // Convert b to float64 for comparison
    floatB, errB := toFloat64(b)
    if errB != nil {
        return nil
    }
    
    // Return nil if values are equal
    if floatA == floatB {
        return nil
    }
    
    // Return the larger value in its original type
    if floatA > floatB {
        return a
    }
    return b
}

// toFloat64 converts various types to float64 for numeric comparison
func toFloat64(v interface{}) (float64, error) {
    switch val := v.(type) {
    case int:
        return float64(val), nil
    case float64:
        return val, nil
    case string:
        // Replace comma with period for European number format
        val = strings.ReplaceAll(val, ",", ".")
        return strconv.ParseFloat(val, 64)
    default:
        return 0, nil
    }
}
