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

    // Filter to keep only integers between 1 and 9 inclusive
    var validNums []int
    for _, v := range arr {
        if v >= 1 && v <= 9 {
            validNums = append(validNums, v)
        }
    }

    // Sort in descending order (reverse=True in Python)
    sort.Slice(validNums, func(i, j int) bool {
        return validNums[i] > validNums[j]
    })

    // Map each valid number to its corresponding name
    var newArr []string
    for _, v := range validNums {
        if name, ok := dic[v]; ok {
            newArr = append(newArr, name)
        }
    }

    return newArr
}

func main() {
    // Test case 1: Normal case
    arr1 := []int{2, 1, 1, 4, 5, 8, 2, 3}
    fmt.Println(ByLength(arr1))

    // Test case 2: Empty array
    arr2 := []int{}
    fmt.Println(ByLength(arr2))

    // Test case 3: With strange numbers
    arr3 := []int{1, -1, 55}
    fmt.Println(ByLength(arr3))
}
