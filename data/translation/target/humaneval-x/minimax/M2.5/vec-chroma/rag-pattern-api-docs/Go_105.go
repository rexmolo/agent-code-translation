package main

import (
    "fmt"
    "sort"
)

func ByLength(arr []int) []string {
    dic := map[int]string{
        1: "One",
        2: "Two",
        3: "Three",
        4: "Four",
        5: "Five",
        6: "Six",
        7: "Seven",
        8: "Eight",
        9: "Nine",
    }

    // Sort in reverse order (descending)
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })

    var newArr []string
    for _, val := range arr {
        if name, ok := dic[val]; ok {
            newArr = append(newArr, name)
        }
    }
    return newArr
}

func main() {
    // Test cases
    arr1 := []int{2, 1, 1, 4, 5, 8, 2, 3}
    fmt.Println(ByLength(arr1)) // [Eight Five Four Three Two Two One One]

    arr2 := []int{}
    fmt.Println(ByLength(arr2)) // []

    arr3 := []int{1, -1, 55}
    fmt.Println(ByLength(arr3)) // [One]
}