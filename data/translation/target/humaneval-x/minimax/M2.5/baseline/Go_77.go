package main

import (
    "math"
)

func Iscube(a int) bool {
    aAbs := int(math.Abs(float64(a)))
    root := math.Pow(float64(aAbs), 1.0/3.0)
    rounded := int(math.Round(root))
    return rounded*rounded*rounded == aAbs
}