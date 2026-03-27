package main

import "math"

// ProdSigns calculates the product of signs multiplied by the sum of absolute values.
// Returns nil for empty arrays.
func ProdSigns(arr []int) interface{} {
    if len(arr) == 0 {
        return nil
    }

    // Check if 0 exists in the array (product of signs becomes 0)
    hasZero := false
    for _, v := range arr {
        if v == 0 {
            hasZero = true
            break
        }
    }

    // Calculate product of signs: 0 if any zero, otherwise (-1)^count(negatives)
    var prod int
    if hasZero {
        prod = 0
    } else {
        negCount := 0
        for _, v := range arr {
            if v < 0 {
                negCount++
            }
        }
        if negCount%2 == 0 {
            prod = 1
        } else {
            prod = -1
        }
    }

    // Calculate sum of absolute values
    sumAbs := 0
    for _, v := range arr {
        sumAbs += int(math.Abs(float64(v)))
    }

    return prod * sumAbs
}

func main() {
    // Test cases from docstring
    // ProdSigns([1, 2, 2, -4]) == -9
    // ProdSigns([0, 1]) == 0
    // ProdSigns([]) == nil
}