package main

func LargestSmallestIntegers(lst []int) [2]interface{} {
    var smallestNeg int
    var hasNeg bool
    var largestPos int
    var hasPos bool

    for _, x := range lst {
        if x < 0 {
            // For negative numbers, we want the maximum (closest to 0)
            if !hasNeg || x > smallestNeg {
                smallestNeg = x
                hasNeg = true
            }
        } else if x > 0 {
            // For positive numbers, we want the minimum
            if !hasPos || x < largestPos {
                largestPos = x
                hasPos = true
            }
        }
    }

    result := [2]interface{}{}
    if hasNeg {
        result[0] = smallestNeg
    }
    if hasPos {
        result[1] = largestPos
    }

    return result
}
