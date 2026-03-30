package main

import "math"

func Iscube(a int) bool {
    if a < 0 {
        a = -a
    }
    cubeRoot := math.Pow(float64(a), 1.0/3.0)
    rounded := int(math.Round(cubeRoot))
    return rounded*rounded*rounded == a
}
