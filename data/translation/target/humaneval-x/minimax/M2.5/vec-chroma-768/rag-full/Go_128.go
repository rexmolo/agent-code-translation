package main

import "fmt"

// ProdSigns calculates the sum of magnitudes multiplied by the product of signs.
// Returns nil for empty array.
func ProdSigns(arr []int) interface{} {
    if len(arr) == 0 {
        return nil
    }

    hasZero := false
    negativeCount := 0
    sumAbs := 0

    for _, v := range arr {
        if v == 0 {
            hasZero = true
        } else {
            if v < 0 {
                negativeCount++
            }
            if v < 0 {
                sumAbs += -v
            } else {
                sumAbs += v
            }
        }
    }

    // If any element is 0, product of signs is 0
    if hasZero {
        return 0
    }

    // Product of signs: (-1)^count of negatives
    var prod int
    if negativeCount%2 == 0 {
        prod = 1
    } else {
        prod = -1
    }

    return prod * sumAbs
}

func main() {
    // Test cases
    fmt.Println(ProdSigns([]int{1, 2, 2, -4})) // -9
    fmt.Println(ProdSigns([]int{0, 1}))        // 0
    fmt.Println(ProdSigns([]int{}))           // nil
}