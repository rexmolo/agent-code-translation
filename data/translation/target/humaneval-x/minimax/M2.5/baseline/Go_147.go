package main

import "fmt"

func GetMaxTriples(n int) int {
    // Count elements by their remainder modulo 3
    // a[i] = i*i - i + 1 = i*(i-1) + 1
    // We only care about val % 3
    count := [3]int{0, 0, 0}

    for i := 1; i <= n; i++ {
        val := i*i - i + 1
        count[val%3]++
    }

    // Calculate valid triples using combinatorics
    // Valid combinations modulo 3: (0,0,0), (1,1,1), (2,2,2), (0,1,2)
    ans := 0

    // (0, 0, 0): C(count[0], 3)
    if count[0] >= 3 {
        ans += count[0] * (count[0]-1) * (count[0]-2) / 6
    }

    // (1, 1, 1): C(count[1], 3)
    if count[1] >= 3 {
        ans += count[1] * (count[1]-1) * (count[1]-2) / 6
    }

    // (2, 2, 2): C(count[2], 3)
    if count[2] >= 3 {
        ans += count[2] * (count[2]-1) * (count[2]-2) / 6
    }

    // (0, 1, 2): count[0] * count[1] * count[2]
    ans += count[0] * count[1] * count[2]

    return ans
}

func main() {
    // Test case
    result := GetMaxTriples(5)
    fmt.Println(result) // Output: 1
}
