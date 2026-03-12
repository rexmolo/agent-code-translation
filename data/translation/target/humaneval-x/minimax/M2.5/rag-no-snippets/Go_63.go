package main

import "bufio"
import "fmt"
import "os"

func Fibfib(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 0
    }
    if n == 2 {
        return 1
    }

    // Use iterative approach for efficiency (O(n) vs exponential time)
    a, b, c := 0, 0, 1 // represents fibfib(n-3), fibfib(n-2), fibfib(n-1)
    for i := 3; i <= n; i++ {
        a, b, c = b, c, a+b+c
    }
    return c
}

func main() {
    in := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscan(in, &n)
    fmt.Println(Fibfib(n))
}
