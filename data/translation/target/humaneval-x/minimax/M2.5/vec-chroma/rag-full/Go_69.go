package main

import (
    "bufio"
    "fmt"
    "os"
)

func Search(lst []int) int {
    // Find maximum element in the list
    maxVal := lst[0]
    for _, v := range lst {
        if v > maxVal {
            maxVal = v
        }
    }

    // Create frequency array
    frq := make([]int, maxVal+1)
    for _, i := range lst {
        frq[i]++
    }

    // Find the greatest integer where frequency >= value
    ans := -1
    for i := 1; i < len(frq); i++ {
        if frq[i] >= i {
            ans = i
        }
    }

    return ans
}

func main() {
    // Read input - assuming format similar to Python examples
    reader := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscan(reader, &n)
    
    lst := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(reader, &lst[i])
    }
    
    result := Search(lst)
    fmt.Println(result)
}
