package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func StrangeSortList(lst []int) []int {
    res := []int{}
    switchFlag := true

    for len(lst) > 0 {
        // Find indices of min and max in a single pass
        minIdx := 0
        maxIdx := 0
        for i, v := range lst {
            if v < lst[minIdx] {
                minIdx = i
            }
            if v > lst[maxIdx] {
                maxIdx = i
            }
        }

        if switchFlag {
            res = append(res, lst[minIdx])
            // Remove element at minIdx
            lst = append(lst[:minIdx], lst[minIdx+1:]...)
        } else {
            res = append(res, lst[maxIdx])
            // Remove element at maxIdx
            lst = append(lst[:maxIdx], lst[maxIdx+1:]...)
        }

        switchFlag = !switchFlag
    }

    return res
}

func main() {
    // Read input from stdin (expected format: space-separated integers)
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter integers separated by spaces: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    var lst []int
    if input != "" {
        parts := strings.Fields(input)
        for _, part := range parts {
            num, _ := strconv.Atoi(part)
            lst = append(lst, num)
        }
    }

    result := StrangeSortList(lst)
    fmt.Println("Result:", result)
}